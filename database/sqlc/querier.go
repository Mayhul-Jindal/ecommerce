// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddOrder(ctx context.Context, arg AddOrderParams) (Order, error)
	AddToCart(ctx context.Context, arg AddToCartParams) (Cart, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeactivateUser(ctx context.Context, id int64) (User, error)
	// TODO: What happens when a book is deleted ?
	DeleteBook(ctx context.Context, id int64) error
	DeleteCartItem(ctx context.Context, arg DeleteCartItemParams) error
	DeleteReview(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) (User, error)
	GetAllTags(ctx context.Context) ([]Tag, error)
	// select b.id, b.title, b.author, b.price, array_agg(r.comment) as comments from "Books" b
	// join "Reviews" r on r.book_id = b.id
	// where b.id = $1
	// group by b.id, b.title, b.author, b.price;
	GetBookById(ctx context.Context, arg GetBookByIdParams) (Book, error)
	GetCartItemsByUserId(ctx context.Context, userID int64) ([]GetCartItemsByUserIdRow, error)
	GetOrderId(ctx context.Context, id int64) (int64, error)
	GetReviewsByBookId(ctx context.Context, arg GetReviewsByBookIdParams) ([]Review, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTotalCartAmountById(ctx context.Context, userID int64) (int64, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
	SearchBookV2(ctx context.Context, arg SearchBookV2Params) ([]SearchBookV2Row, error)
	// TODO fuzzy searching add karni hain isme fkin
	SearchBooksV1(ctx context.Context, arg SearchBooksV1Params) ([]SearchBooksV1Row, error)
	UpdateReview(ctx context.Context, arg UpdateReviewParams) (Review, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
