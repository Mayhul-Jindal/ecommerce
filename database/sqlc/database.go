// TODO
/*
- postgres connection
- creating test for these things
- schema migration
*/

package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storer interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
	UpdateOrderTx(ctx context.Context, arg UpdateOrderTxParams) (UpdateOrderTxResult, error)
	DeleteUserTx(ctx context.Context, userID int64) error
}

// SQLStore provides all functions to execute SQL queries and transactions
type postgresDb struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewPostgresStore(connPool *pgxpool.Pool) Storer {
	return &postgresDb{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// ExecTx executes a function within a database transaction
func (p *postgresDb) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := p.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
