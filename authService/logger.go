package authService

import (
	"context"
	"os"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/rs/zerolog"
)

// logger for auth service
type loggingService struct {
	next   Manager
	logger zerolog.Logger
}

func NewLoggingService(svc Manager) Manager {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &loggingService{
		next:   svc,
		logger: logger,
	}
}

func (l *loggingService) SignUp(ctx context.Context, req types.CreateUserRequest) (res types.UserResponse, err error) {
	defer func(begin time.Time) {
		if err != nil{
			l.logger.Error().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.ID).
			Str("err", err.Error()).
			Dur("took", time.Since(begin)).
			Send()

		} else{
			l.logger.Info().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.ID).
			Dur("took", time.Since(begin)).
			Send()
		}
	}(time.Now())

	return l.next.SignUp(ctx, req)
}


func (l *loggingService) Login(ctx context.Context, req types.LoginUserRequest) (res types.LoginUserResponse, err error) {
	defer func(begin time.Time) {
		if err != nil{
			l.logger.Error().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.User.ID).
			Str("err", err.Error()).
			Dur("took", time.Since(begin)).
			Send()

		} else{
			l.logger.Info().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.User.ID).
			Dur("took", time.Since(begin)).
			Send()
		}
	}(time.Now())

	return l.next.Login(ctx, req)
}

func (l *loggingService) RenewAccess(ctx context.Context, req types.RenewAccessTokenRequest) (res types.RenewAccessTokenResponse, err error) {
	defer func(begin time.Time) {
		if err != nil{
			l.logger.Error().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.UserID).
			Str("err", err.Error()).
			Dur("took", time.Since(begin)).
			Send()

		} else{
			l.logger.Info().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.UserID).
			Dur("took", time.Since(begin)).
			Send()
		}
	}(time.Now())

	return l.next.RenewAccess(ctx, req)
}

func (l *loggingService) VerifyEmail(ctx context.Context, id int64, secret_code string) (res types.VerifyEmailResponse, err error) {
	defer func(begin time.Time) {
		if err != nil{
			l.logger.Error().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.UserID).
			Str("err", err.Error()).
			Dur("took", time.Since(begin)).
			Send()

		} else{
			l.logger.Info().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", res.UserID).
			Dur("took", time.Since(begin)).
			Send()
		}
	}(time.Now())

	return l.next.VerifyEmail(ctx, id, secret_code)
}

func (l *loggingService) ResendEmail(ctx context.Context, req types.ResendEmailRequest) {
	defer func(begin time.Time) {
		l.logger.Info().
		Str("addr", ctx.Value(types.RemoteAddress).(string)).
		Int64("id", req.UserID).
		Dur("took", time.Since(begin)).
		Send()
	}(time.Now())

	l.next.ResendEmail(ctx, req)
}

func (l *loggingService) DeactivateAccount(ctx context.Context, req types.DeactivateAccountRequest) (res types.DeactivateAccountResponse, err error) {
	defer func(begin time.Time) {
		if err != nil{
			l.logger.Error().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", req.UserID).
			Str("err", err.Error()).
			Dur("took", time.Since(begin)).
			Send()

		} else{
			l.logger.Info().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", req.UserID).
			Dur("took", time.Since(begin)).
			Send()
		}
	}(time.Now())

	return l.next.DeactivateAccount(ctx, req)
}

func (l *loggingService) DeleteAccount(ctx context.Context, req types.DeleteAccountRequest) (res types.DeleteAccountResponse, err error) {
	defer func(begin time.Time) {
		if err != nil{
			l.logger.Error().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", req.UserID).
			Str("err", err.Error()).
			Dur("took", time.Since(begin)).
			Send()

		} else{
			l.logger.Info().
			Str("addr", ctx.Value(types.RemoteAddress).(string)).
			Int64("id", req.UserID).
			Dur("took", time.Since(begin)).
			Send()
		}
	}(time.Now())

	return l.next.DeleteAccount(ctx, req)
}