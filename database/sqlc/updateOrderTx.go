package database

import (
	"context"
)

type UpdateOrderTxParams struct {
	UserID  int64
	OrderID int64
}

type UpdateOrderTxResult struct {
	Order     Order
	Purchases []Purchase
}

func (p *postgresDb) UpdateOrderTx(ctx context.Context, arg UpdateOrderTxParams) (UpdateOrderTxResult, error) {
	var result UpdateOrderTxResult

	err := p.execTx(ctx, func(q *Queries) error {
		var err error

		result.Order, err = q.UpdateOrder(ctx, UpdateOrderParams{
			Status: "verified",
			ID:     arg.OrderID,
			UserID: arg.UserID,
		})
		if err != nil {
			return err
		}

		books, err := q.GetCartItemsByUserId(ctx, arg.UserID)
		if err != nil {
			return err
		}

		for _, v := range books {
			purchase, err := q.CreatePurchase(ctx, CreatePurchaseParams{
				UserID:  arg.UserID,
				OrderID: arg.OrderID,
				BookID:  v.ID,
			})
			if err != nil {
				return err
			}

			result.Purchases = append(result.Purchases, purchase)

			err = q.DeleteCartItem(ctx, DeleteCartItemParams{
				UserID: arg.UserID,
				BookID: v.ID,
			})
			if err != nil {
				return err
			}
		}

		return err
	})
	return result, err
}
