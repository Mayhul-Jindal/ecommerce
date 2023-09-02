package database

import (
	"context"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user GetUserParams)
}

type CreateUserTxResult struct {
	User User
}

func (p *postgresDb) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := p.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		arg.AfterCreate(GetUserParams{
			ID:       result.User.ID,
			Username: result.User.Username,
		})

		
		return nil
	})

	return result, err
}
