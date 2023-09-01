// types for book microservice
package types

type SearchBooksV1Request struct {
	WebsearchToTsquery string `json:"websearch_to_tsquery" validate:"required"`
	Limit              int32  `json:"limit" validate:"required,min=1,max=20"`
	Offset             int32  `json:"offset" validate:"min=0"`
}

type GetBookRequest struct {
	UserID int64 `json:"user_id" validate:"required,number,min=1"`
	BookID int64 `json:"book_id" validate:"required,number,min=1"`
}

type AddBookRequest struct {
	UserID       int64   `json:"user_id" validate:"required,number,min=1"`
	Title        string  `json:"title" validate:"required,min=1"`
	Author       string  `json:"author" validate:"required,min=1"`
	TagsArray    []int32 `json:"tags_array"`
	Price        int32   `json:"price" validate:"required,number,min=1"`
	Description  string  `json:"description" validate:"required,min=1"`
	DownloadLink string  `json:"download_link" validate:"required,url"`
}

type UpdateBookRequest struct {
	UserID       int64   `json:"user_id" validate:"required,number,min=1"`
	BookID       int64   `json:"book_id" validate:"required,number,min=1"`
	Title        string  `json:"title" validate:"required,min=1"`
	Author       string  `json:"author" validate:"required,min=1"`
	TagsArray    []int32 `json:"tags_array"`
	Price        int32   `json:"price" validate:"required,number,min=1"`
	Description  string  `json:"description" validate:"required,min=1"`
	DownloadLink string  `json:"download_link" validate:"required,url"`
}

type DeleteBookRequest struct {
	UserID int64 `json:"user_id" validate:"required,number,min=1"`
	BookID int64 `json:"book_id" validate:"required,number,min=1"`
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
	UserID     int64   `json:"user_id" validate:"required,number,min=1"`
	TotalMoney float64 `json:"total_money" validate:"required,number,min=0"`
}

type VerifyOrderRequest struct {
	UserID            int64  `json:"user_id" validate:"required,number,min=1"`
	OrderID           int64  `json:"order_id" validate:"required,number,min=1"`
	RazorpayPaymentID string `json:"razorpay_payment_id" validate:"required"`
	RazorpayOrderID   string `json:"razorpay_order_id" validate:"required"`
	RazorpaySignature string `json:"razorpay_signature" validate:"required"`
}

type GetReviewsRequest struct {
	BookID int64 `json:"book_id" validate:"required,number,min=1"`
}

type AddReviewRequest struct {
	UserID  int64  `json:"user_id" validate:"required,number,min=1"`
	BookID  int64  `json:"book_id" validate:"required,number,min=1"`
	Rating  int32  `json:"rating" valideate:"required,number,min=1"`
	Comment string `json:"comment" valideate:"required,min=5"`
}

type UpdateReviewRequest struct {
	UserID  int64  `json:"user_id" validate:"required,number,min=1"`
	ID      int64  `json:"id" validate:"required,number,min=1"`
	Rating  int32  `json:"rating" valideate:"required,number,min=1"`
	Comment string `json:"comment" valideate:"required,min=5"`
}

type DeleteReviewRequest struct {
	ID int64 `json:"id" validate:"required,number,min=1"`
}

type GetHotSellingRequest struct {
	Limit  int32 `json:"limit" validate:"required,number,min=1"`
	Offset int32 `json:"offset" validate:"required,number,min=0"`
}

type GetPersonalRecommendationsRequest struct {
	UserID int64 `json:"user_id" validate:"required,number,min=1"`
}
