package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/email"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
)

type Worker interface {
	EnqueueSendVerifyEmail(req database.GetUserParams) 
	EnqueueDeleteOperation(req database.GetUserParams)
}

type worker struct {
	emailCh chan struct{}
	delCh   chan struct{}
	quitCh  chan struct{}
	db      database.Storer
	email   email.EmailSender
}

func NewWorker(db database.Storer, email email.EmailSender) Worker {
	return &worker{
		// this means maximum of 20 emails can be processed at once
		emailCh: make(chan struct{}, 20),
		quitCh:  make(chan struct{}),
		db:      db,
		email:   email,
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
		log.Printf("msg=timeout\ttook=%v", time.Since(start))
		return

	case <-w.quitCh:
		log.Printf("msg=quit\ttook=%v", time.Since(start))
		return

	case msg := <-msgCh:
		log.Printf("msg=%v\ttook=%v", msg, time.Since(start))
		return
	}
}






func (w *worker) EnqueueDeleteOperation(req database.GetUserParams) {
	w.delCh <- struct{}{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func() {
		w.deleteOperation(ctx, req)
		cancel()
		<-w.delCh
	}()
}

func (w *worker) deleteOperation(tx context.Context, req database.GetUserParams) {
}


