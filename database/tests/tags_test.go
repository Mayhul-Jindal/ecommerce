package database_test

import (
	"context"
	"testing"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTag(t *testing.T) database.Tag {
	ranTag := database.CreateTagParams{
		ID: util.RandomInt(1, 10),
		TagName: util.RandomString(10),
	}  

	tag, err := testQueries.CreateTag(context.Background(), ranTag)
	require.NoError(t, err)
	require.NotEmpty(t, tag)

	require.Equal(t, tag.TagName, ranTag.TagName)

	return tag
}

func TestCreateTag(t *testing.T) {
	CreateRandomTag(t)
}

