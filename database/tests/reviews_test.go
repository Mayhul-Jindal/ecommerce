package sqlc_test

import (
	"context"
	"testing"

	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) sqlc.User{
	return sqlc.User{}
}

func CreateRandomReview(t *testing.T) sqlc.Review{
	review := sqlc.CreateReviewParams{
		UserID: util.RandomBigInt(1, 100),
		BookID: util.RandomBigInt(1, 100),
		Rating: util.RandomInt(1, 5),
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
