package database_test

import (
	"context"
	"testing"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) database.User {
	return database.User{}
}

func CreateRandomReview(t *testing.T) database.Review {
	review := database.CreateReviewParams{
		UserID:  util.RandomBigInt(1, 50),
		BookID:  util.RandomBigInt(51, 1000),
		Rating:  util.RandomInt(1, 5),
		Comment: util.RandomString(10),
	}

	review1, err := testQueries.CreateReview(context.Background(), review)
	require.NoError(t, err)
	require.NotEmpty(t, review1)

	require.Equal(t, review, review1)

	return review1
}

func TestCreateReview(t *testing.T) {
	CreateRandomReview(t)
}
