package database

import (
	"context"
)

// purchases --> carts --> orders --> verify_email --> sessions --> users
func (p *postgresDb) DeleteUserTx(ctx context.Context, userID int64) error {
	err := p.execTx(ctx, func(q *Queries) error {
		var err error

		err = p.DeletePurchasesOfUser(ctx, userID)
		if err != nil {
			return err
		}

		err = p.DeleteCartOfUser(ctx, userID)
		if err != nil {
			return err
		}

		err = p.DeleteOrdersOfUser(ctx, userID)
		if err != nil {
			return err
		}

		err = p.DeleteVerifyEmailsOfUser(ctx, userID)
		if err != nil {
			return err
		}

		err = p.DeleteSessionsOfUser(ctx, userID)
		if err != nil {
			return err
		}

		err = p.DeleteUser(ctx, userID)
		if err != nil {
			return err
		}

		return err
	})

	return err
}
