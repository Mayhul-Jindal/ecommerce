// TODO
/*
- make error.go jisme one should put all types of errors in sqlc package
*/

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

// helper function for my tests. this is to avoid dependencies between tests (each test sare independent)
func CreateRandomBook(t *testing.T) Book {
	book1 := CreateBookParams{
		Title:       util.RandomString(10),
		Author:      util.RandomString(10),
		Tags:        util.RandomTags(),
		Price:       util.RandomInt(1, 1000),
		Quantity:    util.RandomInt(1, 10),
		Description: util.RandomString(50),
	}

	book2, err := testQueries.CreateBook(context.Background(), book1)
	// basic checks
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	// equality checks
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Tags, book2.Tags)
	require.Equal(t, book1.Price, book2.Price)
	require.Equal(t, book1.Quantity, book2.Quantity)
	require.Equal(t, book1.Description, book2.Description)

	// book2 checks
	require.NotZero(t, book2.ID)
	require.NotZero(t, book2.CreatedAt)
	return book2
}

func TestCreateBook(t *testing.T) {
	CreateRandomBook(t)
}

func TestGetBook(t *testing.T) {
	book1 := CreateRandomBook(t)
	book2, err := testQueries.GetBook(context.Background(), book1.ID)
	// basic checks
	require.NoError(t, err)
	require.NotEmpty(t, book2)

	// equality checks
	require.Equal(t, book1.ID, book2.ID)
	require.Equal(t, book1.Title, book2.Title)
	require.Equal(t, book1.Author, book2.Author)
	require.Equal(t, book1.Tags, book2.Tags)
	require.Equal(t, book1.Price, book2.Price)
	require.Equal(t, book1.Quantity, book2.Quantity)
	require.Equal(t, book1.Description, book2.Description)
	require.WithinDuration(t, book1.CreatedAt, book2.CreatedAt, time.Second)
}

func TestGetBooks(t *testing.T) {
	var lastBook Book
	for i := 0; i < 5; i++ {
		lastBook = CreateRandomBook(t)
	}

	arg := GetBooksParams {
		Author: lastBook.Author,
		Limit: 5,
		Offset: 0,
	}
	books, err := testQueries.GetBooks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, books)

	for _, book := range books {
		require.NotEmpty(t, book)
		require.Equal(t, lastBook.Author, book.Author)
	}
}	

func TestUpdateBookDesc(t *testing.T) {
	book1 := CreateRandomBook(t)

	newDesc := util.RandomString(10)
	arg := UpdateBookDescParams{
		ID: book1.ID,
		Description: newDesc,
	}

	Book2, err := testQueries.UpdateBookDesc(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Book2)

	require.Equal(t, book1.ID, Book2.ID)
	require.Equal(t, book1.Title, Book2.Title)
	require.Equal(t, Book2.Description, newDesc)
	require.WithinDuration(t, book1.CreatedAt, Book2.CreatedAt, time.Second)
}

func TestDeleteBook(t *testing.T) {
	book1 := CreateRandomBook(t)
	err := testQueries.DeleteBook(context.Background(), book1.ID)
	require.NoError(t, err)

	book2, err := testQueries.GetBook(context.Background(), book1.ID)
	require.Error(t, err)
	require.Empty(t, book2)
}
