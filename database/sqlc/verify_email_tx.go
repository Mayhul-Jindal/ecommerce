package database

import (
	"context"

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

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.ID,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = p.UpdateUser(ctx, UpdateUserParams{
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
