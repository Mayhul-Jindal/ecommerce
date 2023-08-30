// types for book microservice
package types

type SearchBooksV1Request struct {
	WebsearchToTsquery string `json:"websearch_to_tsquery" validate:"required"`
	Limit              int32  `json:"limit" validate:"required,min=1,max=20"`
	Offset             int32  `json:"offset" validate:"min=0"`
}

type GetCartRequest struct {
	UserID int64 `json:"user_id" validate:"required,number,min=1"`
}

type AddToCartRequest struct {
	UserID int64 `json:"user_id" validate:"required,number,min=1"`
	BookID int64 `json:"book_id" validate:"required,number,min=1"`
}

type DeleteCartItemRequest struct {
	UserID int64 `json:"user_id" validate:"required,number,min=1"`
	BookID int64 `json:"book_id" validate:"required,number,min=1"`
}

type PlaceOrderRequest struct {
	UserID     int64   `json:"user_id" validate:"required"`
	TotalMoney float64 `json:"total_money" validate:"required,number,min=0"`
}
