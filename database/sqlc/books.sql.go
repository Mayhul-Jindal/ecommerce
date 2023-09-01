// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: books.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBook = `-- name: CreateBook :one
INSERT INTO "Books" (
  title, author, tags_array, price, description, download_link
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, title, author, tags_array, price, description, download_link, created_at
`

type CreateBookParams struct {
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	TagsArray    []int32 `json:"tags_array"`
	Price        int32   `json:"price"`
	Description  string  `json:"description"`
	DownloadLink string  `json:"download_link"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, createBook,
		arg.Title,
		arg.Author,
		arg.TagsArray,
		arg.Price,
		arg.Description,
		arg.DownloadLink,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.TagsArray,
		&i.Price,
		&i.Description,
		&i.DownloadLink,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM "Books"
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteBook, id)
	return err
}

const getBookById = `-- name: GetBookById :one
select title, author, price, description, download_link from "Books"
where id = $1
`

type GetBookByIdRow struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Price        int32  `json:"price"`
	Description  string `json:"description"`
	DownloadLink string `json:"download_link"`
}

func (q *Queries) GetBookById(ctx context.Context, id int64) (GetBookByIdRow, error) {
	row := q.db.QueryRow(ctx, getBookById, id)
	var i GetBookByIdRow
	err := row.Scan(
		&i.Title,
		&i.Author,
		&i.Price,
		&i.Description,
		&i.DownloadLink,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :one
UPDATE "Books"
SET
  title = COALESCE($1, title),
  author = COALESCE($2, author),
  tags_array = COALESCE($3, tags_array),
  price = COALESCE($4, price),
  description = COALESCE($5, description),
  download_link = COALESCE($6, download_link)
WHERE
  id = $7
RETURNING id, title, author, tags_array, price, description, download_link, created_at
`

type UpdateBookParams struct {
	Title        pgtype.Text `json:"title"`
	Author       pgtype.Text `json:"author"`
	TagsArray    []int32     `json:"tags_array"`
	Price        pgtype.Int4 `json:"price"`
	Description  pgtype.Text `json:"description"`
	DownloadLink pgtype.Text `json:"download_link"`
	ID           int64       `json:"id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, updateBook,
		arg.Title,
		arg.Author,
		arg.TagsArray,
		arg.Price,
		arg.Description,
		arg.DownloadLink,
		arg.ID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.TagsArray,
		&i.Price,
		&i.Description,
		&i.DownloadLink,
		&i.CreatedAt,
	)
	return i, err
}
