// TODO
/*
- database integration
- make these endpoints work with test
*/

package bookService

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/token"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/types"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/razorpay/razorpay-go"
)

type Manager interface {
	// search
	Search(ctx context.Context, req types.SearchBooksV1Request) ([]database.SearchBooksV2Row, error)

	// recommendations
	GetHotSelling(ctx context.Context, req types.GetHotSellingRequest) ([]database.GetHotSellingBooksRow, error)
	GetPersonalRecommendations(ctx context.Context, req types.GetPersonalRecommendationsRequest) ([]database.SearchBooksV2Row, error)

	// books
	GetBook(ctx context.Context, req types.GetBookRequest) (database.GetBookByIdRow, error)
	AddBook(ctx context.Context, req types.AddBookRequest) (database.Book, error)
	UpdateBook(ctx context.Context, req types.UpdateBookRequest) (database.Book, error)

	// tags
	GetAllTags(ctx context.Context) ([]database.Tag, error)
	CreateTag(ctx context.Context, req types.CreateTagRequest) (database.Tag, error)
	UpdateTag(ctx context.Context, req types.UpdateTagRequest) (database.Tag, error)
	DeleteTag(ctx context.Context, req types.DeleteTagRequest) (error)

	// cart
	GetCart(ctx context.Context, req types.GetCartRequest) ([]database.GetCartItemsByUserIdRow, error)
	AddToCart(ctx context.Context, req types.AddToCartRequest) (database.Cart, error)
	DeleteCartItem(ctx context.Context, req types.DeleteCartItemRequest) error

	// place order
	PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (database.Order, error)
	VerifyOrder(ctx context.Context, req types.VerifyOrderRequest) (database.UpdateOrderTxResult, error)
	GetPurchases(ctx context.Context, req types.GetPurchasesRequest) ([]database.GetPurchasedBooksRow, error)

	// reviews
	GetReviews(ctx context.Context, req types.GetReviewsRequest) ([]database.Review, error)
	AddReview(ctx context.Context, req types.AddReviewRequest) (database.Review, error)
	UpdateReview(ctx context.Context, req types.UpdateReviewRequest) (database.Review, error)
	DeleteReview(ctx context.Context, req types.DeleteReviewRequest) error
}

// one should add dependencies/tools here
type bookManager struct {
	db             database.Storer
	razorPayClient *razorpay.Client
}

func NewManager(db database.Storer, razorPayClient *razorpay.Client) Manager {
	return &bookManager{
		db:             db,
		razorPayClient: razorPayClient,
	}
}

func (b *bookManager) Search(ctx context.Context, req types.SearchBooksV1Request) ([]database.SearchBooksV2Row, error) {
	params := database.SearchBooksV2Params{
		WebsearchToTsquery: convertToSearchString(req.WebsearchToTsquery),
		Difference:         req.WebsearchToTsquery,
		Limit:              req.Limit,
		Offset:             req.Offset,
	}

	resp, err := b.db.SearchBooksV2(ctx, params)
	if err != nil {
		return []database.SearchBooksV2Row{}, err
	}

	return resp, nil
}

func (b *bookManager) GetHotSelling(ctx context.Context, req types.GetHotSellingRequest) ([]database.GetHotSellingBooksRow, error) {
	resp, err := b.db.GetHotSellingBooks(ctx, database.GetHotSellingBooksParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return []database.GetHotSellingBooksRow{}, err
	}

	return resp, nil
}

func (b *bookManager) GetPersonalRecommendations(ctx context.Context, req types.GetPersonalRecommendationsRequest) ([]database.SearchBooksV2Row, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return nil, errs.ErrorNotAuthorized
	}

	resp, err := b.db.GetUserRecommendations(ctx, database.GetUserRecommendationsParams{
		UserID:  authPayload.UserID,
		OrderID: req.OrderID,
	})
	if err != nil {
		return nil, err
	}

	var str string
	for _, b := range resp {
		str += string(b) + " "
	}

	log.Println(str)
	recom, err := b.db.SearchBooksV2(ctx, database.SearchBooksV2Params{
		WebsearchToTsquery: convertToSearchString(string(str)),
		Limit: req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}

	return recom, nil
}

// authorized
func (b *bookManager) GetBook(ctx context.Context, req types.GetBookRequest) (database.GetBookByIdRow, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.GetBookByIdRow{}, errs.ErrorNotAuthorized
	}

	resp, err := b.db.GetBookById(ctx, req.BookID)
	if err != nil {
		return database.GetBookByIdRow{}, err
	}

	_, err = b.db.CheckBookPurchased(ctx, database.CheckBookPurchasedParams{
		UserID: authPayload.UserID,
		BookID: req.BookID,
	})
	if err != nil {
		resp.DownloadLink = ""
		return resp, nil
	}

	return resp, nil
}

// admin only routes
func (b *bookManager) AddBook(ctx context.Context, req types.AddBookRequest) (database.Book, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Book{}, errs.ErrorNotAuthorized
	}

	_, err := b.db.CheckAdmin(ctx, authPayload.UserID)
	if err != nil {
		return database.Book{}, errs.ErrorNotAuthorized
	}

	resp, err := b.db.CreateBook(ctx, database.CreateBookParams{
		Title:        req.Title,
		Author:       req.Author,
		TagsArray:    req.TagsArray,
		Price:        req.Price,
		Description:  req.Description,
		DownloadLink: req.DownloadLink,
	})
	if err != nil {
		return database.Book{}, err
	}

	return resp, nil
}

func (b *bookManager) UpdateBook(ctx context.Context, req types.UpdateBookRequest) (database.Book, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Book{}, errs.ErrorNotAuthorized
	}

	_, err := b.db.CheckAdmin(ctx, authPayload.UserID)
	if err != nil {
		return database.Book{}, errs.ErrorNotAuthorized
	}

	resp, err := b.db.UpdateBook(ctx, database.UpdateBookParams{
		ID: req.BookID,
		Title: pgtype.Text{
			String: req.Title,
			Valid:  true,
		},
		Author: pgtype.Text{
			String: req.Author,
			Valid:  true,
		},
		TagsArray: req.TagsArray,
		Price: pgtype.Int4{
			Int32: req.Price,
			Valid: true,
		},
		Description: pgtype.Text{
			String: req.Description,
			Valid:  true,
		},
		DownloadLink: pgtype.Text{
			String: req.DownloadLink,
			Valid:  true,
		},
	})
	if err != nil {
		return database.Book{}, err
	}

	return resp, nil
}

// func (b *bookManager) DeleteBook(ctx context.Context, req types.DeleteBookRequest) (error) {
// 	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
// 	if authPayload.UserID != req.UserID {
// 		return database.Book{}, errs.ErrorNotAuthorized
// 	}

// 	err := b.db.CheckAdmin(ctx, authPayload.UserID)
// 	if err != nil {
// 		return database.Book{}, errs.ErrorNotAuthorized
// 	}
// }

func (b *bookManager) GetAllTags(ctx context.Context) ([]database.Tag, error) {
	resp, err := b.db.GetAllTags(ctx)
	if err != nil {
		return []database.Tag{}, err
	}

	return resp, nil
}


func (b *bookManager) CreateTag(ctx context.Context, req types.CreateTagRequest) (database.Tag, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Tag{}, errs.ErrorNotAuthorized
	}

	_, err := b.db.CheckAdmin(ctx, authPayload.UserID)
	if err != nil {
		return database.Tag{}, errs.ErrorNotAuthorized
	}

	tag, err := b.db.CreateTag(ctx, database.CreateTagParams{
		ID: req.TagID,
		TagName: req.TagName,
	})
	if err != nil {
		return database.Tag{}, err
	}

	return tag, nil
}

func (b *bookManager) UpdateTag(ctx context.Context, req types.UpdateTagRequest) (database.Tag, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Tag{}, errs.ErrorNotAuthorized
	}

	_, err := b.db.CheckAdmin(ctx, authPayload.UserID)
	if err != nil {
		return database.Tag{}, errs.ErrorNotAuthorized
	}

	tag, err := b.db.UpdateTag(ctx, database.UpdateTagParams{
		ID: req.TagID,
		TagName: req.TagName,
	})
	if err != nil {
		return database.Tag{}, err
	}

	return tag, nil
}

func (b *bookManager) DeleteTag(ctx context.Context, req types.DeleteTagRequest) (error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return errs.ErrorNotAuthorized
	}

	_, err := b.db.CheckAdmin(ctx, authPayload.UserID)
	if err != nil {
		return errs.ErrorNotAuthorized
	}

	err = b.db.DeleteTag(ctx, req.TagID)
	if err != nil {
		return err
	}

	return nil
}





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
	// todo check if already bought then dont put in cart
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Cart{}, errs.ErrorNotAuthorized
	}

	_, err := b.db.CheckBookPurchased(ctx, database.CheckBookPurchasedParams{
		UserID: req.UserID,
		BookID: req.BookID,
	})
	if err == nil {
		return database.Cart{}, errs.ErrorBookAlreadyBought
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

func (b *bookManager) PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (database.Order, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Order{}, errs.ErrorNotAuthorized
	}

	// check if email is verified or not
	_, err := b.db.CheckEmailVerified(ctx, authPayload.UserID)
	if err != nil {
		return database.Order{}, errs.ErrorEmailNotVerified
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
		UserID:          authPayload.UserID,
		RazorpayOrderID: razorPayResp["id"].(string),
		TotalMoney:      razorPayResp["amount"].(float64),
		Status:          "pending",
	}

	resp, err := b.db.AddOrder(ctx, params)
	if err != nil {
		return database.Order{}, err
	}

	return resp, nil
}

func (b *bookManager) VerifyOrder(ctx context.Context, req types.VerifyOrderRequest) (database.UpdateOrderTxResult, error) {
	// todo check if already verified
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.UpdateOrderTxResult{}, errs.ErrorNotAuthorized
	}

	// here this is order_id of my database
	order, err := b.db.GetOrderById(ctx, database.GetOrderByIdParams{
		UserID: authPayload.UserID,
		ID:     req.OrderID,
	})
	if err != nil {
		return database.UpdateOrderTxResult{}, err
	}

	params := map[string]interface{}{
		"razorpay_order_id":   order.RazorpayOrderID,
		"razorpay_payment_id": req.RazorpayPaymentID,
	}

	if !VerifyPaymentSignature(params, req.RazorpaySignature, b.razorPayClient.Account.Request.Auth.Secret) {
		return database.UpdateOrderTxResult{}, errs.ErrorPayementNotVerified
	}

	// paymeny verify hogayi now here cpmes the transaction {update order status, add books}
	resp, err := b.db.UpdateOrderTx(ctx, database.UpdateOrderTxParams{
		UserID:  authPayload.UserID,
		OrderID: req.OrderID,
	})
	if err != nil {
		return database.UpdateOrderTxResult{}, err
	}

	return resp, nil
}

func (b *bookManager) GetPurchases(ctx context.Context, req types.GetPurchasesRequest) ([]database.GetPurchasedBooksRow, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return []database.GetPurchasedBooksRow{}, errs.ErrorNotAuthorized
	}

	resp, err := b.db.GetPurchasedBooks(ctx, authPayload.UserID)
	if err != nil {
		return []database.GetPurchasedBooksRow{}, err
	}

	return resp, nil
}

func (b *bookManager) GetReviews(ctx context.Context, req types.GetReviewsRequest) ([]database.Review, error) {
	reviews, err := b.db.GetReviewsByBookId(ctx, database.GetReviewsByBookIdParams{
		BookID: req.BookID,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return []database.Review{}, err
	}

	return reviews, nil
}

func (b *bookManager) AddReview(ctx context.Context, req types.AddReviewRequest) (database.Review, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Review{}, errs.ErrorNotAuthorized
	}

	review, err := b.db.CreateReview(ctx, database.CreateReviewParams{
		UserID:  authPayload.UserID,
		BookID:  req.BookID,
		Rating:  req.Rating,
		Comment: req.Comment,
	})
	if err != nil {
		return database.Review{}, err
	}

	return review, nil
}

func (b *bookManager) UpdateReview(ctx context.Context, req types.UpdateReviewRequest) (database.Review, error) {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return database.Review{}, errs.ErrorNotAuthorized
	}

	review, err := b.db.UpdateReview(ctx, database.UpdateReviewParams{
		ID: req.ReviewID,
		Rating: pgtype.Int4{
			Int32: req.Rating,
			Valid: true,
		},
		Comment: pgtype.Text{
			String: req.Comment,
			Valid:  true,
		},
	})
	if err != nil {
		return database.Review{}, err
	}

	return review, nil
}

func (b *bookManager) DeleteReview(ctx context.Context, req types.DeleteReviewRequest) error {
	authPayload := ctx.Value(types.AuthorizationPayload).(*token.Payload)
	if authPayload.UserID != req.UserID {
		return errs.ErrorNotAuthorized
	}

	err := b.db.DeleteReview(ctx, req.ReviewtID)
	if err != nil {
		return err
	}

	return nil
}

// helper function
func convertToSearchString(sentence string) string {
	words := strings.Split(sentence, " ")
	modifiedSentence := strings.Join(words, " or ")
	return modifiedSentence
}

func VerifyPaymentSignature(queryParams map[string]interface{}, webhookSignature string, webhookSecret string) bool {
	payload := fmt.Sprint(queryParams["razorpay_order_id"], "|", queryParams["razorpay_payment_id"])

	isValid := VerifySignature([]byte(payload), webhookSignature, webhookSecret)

	return isValid
}

func VerifySignature(body []byte, signature string, key string) bool {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return expectedSignature == signature
}
