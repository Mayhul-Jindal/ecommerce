// TODO
/*
- make error.go jisme one should put all types of errors in database package
*/

package database_test

import (
	"context"
	"testing"
	"time"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

// helper function for my tests. this is to avoid dependencies between tests (each test sare independent)
func CreateRandomBook(t *testing.T) database.Book {
	book1 := database.CreateBookParams{
		Title:       util.RandomString(10),
		Author:      util.RandomString(10),
		TagsArray:   util.RandomIntArray(1, 5),
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
	require.Equal(t, book1.TagsArray, book2.TagsArray)
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
	require.Equal(t, book1.TagsArray, book2.TagsArray)
	require.Equal(t, book1.Price, book2.Price)
	require.Equal(t, book1.Quantity, book2.Quantity)
	require.Equal(t, book1.Description, book2.Description)
	require.WithinDuration(t, book1.CreatedAt, book2.CreatedAt, time.Second)
}

// this test is bit buggy
func TestGetBooks(t *testing.T) {
	var lastBook database.Book
	for i := 0; i < 5; i++ {
		lastBook = CreateRandomBook(t)
	}

	arg := database.GetBooksParams{
		Limit:  5,
		Offset: 0,
	}
	books, err := testQueries.GetBooks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, books)

	require.Equal(t, lastBook.Author, books[len(books)-1].Author)
}

func TestUpdateBookDesc(t *testing.T) {
	book1 := CreateRandomBook(t)

	newDesc := util.RandomString(10)
	arg := database.UpdateBookDescParams{
		ID:          book1.ID,
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
