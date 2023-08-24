// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package postgres

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	DeleteBook(ctx context.Context, id int64) error
	GetBook(ctx context.Context, id int64) (Book, error)
	GetBooks(ctx context.Context, arg GetBooksParams) ([]Book, error)
	UpdateBookDesc(ctx context.Context, arg UpdateBookDescParams) (Book, error)
}

var _ Querier = (*Queries)(nil)
