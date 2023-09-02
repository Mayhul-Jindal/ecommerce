package database

import (
	"context"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/jackc/pgx/v5/pgtype"
)

type VerifyEmailTxParams struct {
	ID         int64
	SecretCode string
}

type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

func (p *postgresDb) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := p.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams(arg))
		if err != nil {
			return errs.ErrorLinkExpired
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			ID: result.VerifyEmail.UserID,
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		})

		return err
	})

	return result, err
}
