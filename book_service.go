// TODO
/*
- database integration
- make these endpoints work with test
*/

package main

import (
	"context"
	"log"
	"strings"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/razorpay/razorpay-go"
)

type BookManager interface {
	// unauthorized
	Search(ctx context.Context, req types.SearchBooksV1Request) ([]database.SearchBooksV1Row, error)

	// authorized
	GetCart(ctx context.Context, req types.GetCartRequest) ([]database.GetCartItemsByUserIdRow, error)
	AddToCart(ctx context.Context, req types.AddToCartRequest) (database.Cart, error)
	DeleteCartItem(ctx context.Context, req types.DeleteCartItemRequest) error

	PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (database.Order, error)

	// DeactivateAccount(ctx context.Context, req types.DeactivateAccountRequest) (database.Order, error)
}

// one should add dependencies/tools here
type bookManager struct {
	db             database.Storer
	razorPayClient *razorpay.Client
}

func NewBookManager(db database.Storer, razorPayClient *razorpay.Client) BookManager {
	return &bookManager{
		db:             db,
		razorPayClient: razorPayClient,
	}
}

func (b *bookManager) Search(ctx context.Context, req types.SearchBooksV1Request) ([]database.SearchBooksV1Row, error) {
	params := database.SearchBooksV1Params{
		WebsearchToTsquery: convertToSearchString(req.WebsearchToTsquery),
		Limit:              req.Limit,
		Offset:             req.Offset,
	}

	resp, err := b.db.SearchBooksV1(ctx, params)
	if err != nil {
		return []database.SearchBooksV1Row{}, err
	}

	return resp, nil
}

// has to authorized
func (b *bookManager) GetCart(ctx context.Context, req types.GetCartRequest) ([]database.GetCartItemsByUserIdRow, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return []database.GetCartItemsByUserIdRow{}, errs.ErrorNotAuthorized
	}

	resp, err := b.db.GetCartItemsByUserId(ctx, authPayload.UserID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *bookManager) AddToCart(ctx context.Context, req types.AddToCartRequest) (database.Cart, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Cart{}, errs.ErrorNotAuthorized
	}

	params := database.AddToCartParams{
		UserID: req.UserID,
		BookID: req.BookID,
	}

	resp, err := b.db.AddToCart(ctx, params)
	if err != nil {
		return database.Cart{}, err
	}

	return resp, nil
}

// TODO dont knpw why but no error is comming
func (b *bookManager) DeleteCartItem(ctx context.Context, req types.DeleteCartItemRequest) error {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return errs.ErrorNotAuthorized
	}

	params := database.DeleteCartItemParams{
		UserID: req.UserID,
		BookID: req.BookID,
	}

	return b.db.DeleteCartItem(ctx, params)
}

// TODO order lines, change schema id hata doh
func (b *bookManager) PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (database.Order, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Order{}, errs.ErrorNotAuthorized
	}

	// step 1: check amount before prceeding
	actualCost, err := b.db.GetTotalCartAmountById(ctx, authPayload.UserID)
	if err != nil {
		return database.Order{}, err
	}

	if req.TotalMoney != float64(actualCost) {
		return database.Order{}, errs.ErrorAmountMismatch
	}

	// step 2: get order ID from razor pay
	data := map[string]interface{}{
		"amount":   actualCost * 100,
		"currency": "INR",
	}

	razorPayResp, err := b.razorPayClient.Order.Create(data, nil)
	if err != nil {
		log.Println(1)
		return database.Order{}, err
	}

	// step 3: update this order id and return the result
	params := database.AddOrderParams{
		RazorpayOrderID: razorPayResp["id"].(string),
		UserID:          authPayload.UserID,
		TotalMoney:      razorPayResp["amount"].(float64),
		Status:          "pending",
	}

	resp, err := b.db.AddOrder(ctx, params)
	if err != nil {
		return database.Order{}, err
	}

	return resp, nil
}

// helper function
func convertToSearchString(sentence string) string {
	words := strings.Split(sentence, " ")
	modifiedSentence := strings.Join(words, " or ")
	return modifiedSentence
}
