// TODO
/*
- integrate zap/zerolog library into this for log rotation
- Database ke liye shayad alag se banega
*/

package main

import (
	"context"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
)

/*
This basically demonstrates how the logging service can be seperated from the bussiness
logic which can therefore increase maintainability.
*/
type loggingService struct {
	next BookManager
}

func NewLoggingService(svc BookManager) BookManager {
	return &loggingService{
		next: svc,
	}
}

func (l *loggingService) Search(ctx context.Context, req types.SearchBooksV1Request) ([]database.SearchBooksV1Row, error) {
	return l.next.Search(ctx, req)
}

func (l *loggingService) GetCart(ctx context.Context, req types.GetCartRequest) ([]database.GetCartItemsByUserIdRow, error) {
	return l.next.GetCart(ctx, req)
}

func (l *loggingService) AddToCart(ctx context.Context, req types.AddToCartRequest) (database.Cart, error) {
	return l.next.AddToCart(ctx, req)
}

func (l *loggingService) DeleteCartItem(ctx context.Context, req types.DeleteCartItemRequest) error {
	return l.next.DeleteCartItem(ctx, req)
}

func (l *loggingService) PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (database.Order, error) {
	return l.next.PlaceOrder(ctx, req)
}