package worker

import (
	"context"
	"fmt"
	"os"
	"time"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/email"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/rs/zerolog"
)

type Worker interface {
	EnqueueSendVerifyEmail(req database.GetUserParams)
	EnqueueDeleteOperation(userID int64)
}

type worker struct {
	emailCh chan struct{}
	delCh   chan struct{}
	quitCh  chan struct{}
	db      database.Storer
	email   email.EmailSender
	logger  zerolog.Logger
}

func NewWorker(db database.Storer, email email.EmailSender) Worker {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &worker{
		// this means maximum of 20 emails can be processed at once
		emailCh: make(chan struct{}, 20),
		delCh:   make(chan struct{}, 20),
		quitCh:  make(chan struct{}),
		db:      db,
		email:   email,
		logger:  logger,
	}
}

func (w *worker) EnqueueSendVerifyEmail(req database.GetUserParams) {
	w.emailCh <- struct{}{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func() {
		w.sendVerifyEmail(ctx, req)
		cancel()
		<-w.emailCh
	}()
}

func (w *worker) sendVerifyEmail(ctx context.Context, req database.GetUserParams) {
	var (
		start = time.Now()
		msgCh = make(chan string)
	)

	go func() {
		user, err := w.db.GetUser(ctx, req)
		if err != nil {
			msgCh <- err.Error()
			return
		}

		verifyEmail, err := w.db.CreateVerifyEmail(ctx, database.CreateVerifyEmailParams{
			UserID:     user.ID,
			Email:      user.Email,
			SecretCode: util.RandomString(32),
		})
		if err != nil {
			msgCh <- err.Error()
			return
		}

		subject := "Welcome to Simple Bank"
		verifyUrl := fmt.Sprintf("http://localhost:3000/users/verify_email?id=%d&secret_code=%s", verifyEmail.ID, verifyEmail.SecretCode)
		content := fmt.Sprintf(
			`Hello %s,<br/>
				Thank you for registering with us!<br/>
				Please <a href="%s">click here</a> to verify your email address.<br/>
				`, user.Username, verifyUrl)
		to := []string{user.Email}

		err = w.email.SendEmail(subject, content, to, nil, nil, nil)
		if err != nil {
			msgCh <- err.Error()
			return
		}

		msgCh <- "ok"
	}()

	select {
	case <-ctx.Done():
		w.logger.Error().
			Str("err", "timout").
			Dur("took", time.Since(start)).
			Send()

		return

	case <-w.quitCh:
		w.logger.Error().
			Str("err", "quit").
			Dur("took", time.Since(start)).
			Send()

		return

	case msg := <-msgCh:
		w.logger.Info().
			Str("msg", msg).
			Dur("took", time.Since(start)).
			Send()

		return
	}
}

func (w *worker) EnqueueDeleteOperation(userID int64) {
	w.delCh <- struct{}{}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	go func() {
		// given a delay here to simulate that after some time this job will run not immediately (data retention)
		time.Sleep(20 * time.Second)
		w.deleteOperation(ctx, userID)
		cancel()
		<-w.delCh
	}()
}

func (w *worker) deleteOperation(ctx context.Context, userID int64) {
	var (
		start = time.Now()
		msgCh = make(chan string)
	)

	go func() {
		err := w.db.DeleteUserTx(ctx, userID)
		if err != nil {
			msgCh <- err.Error()
			return
		}

		msgCh <- "ok"
	}()

	select {
	case <-ctx.Done():
		w.logger.Error().
			Str("err", "timout").
			Dur("took", time.Since(start)).
			Send()

		return

	case <-w.quitCh:
		w.logger.Error().
			Str("err", "quit").
			Dur("took", time.Since(start)).
			Send()

		return

	case msg := <-msgCh:
		w.logger.Info().
			Str("msg", msg).
			Dur("took", time.Since(start)).
			Send()

		return
	}
}
