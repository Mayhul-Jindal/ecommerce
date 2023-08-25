package database_test

import (
	"context"
	"testing"

	database "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/database/sqlc"
	"github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTag(t *testing.T) database.Tag {
	ranTag := util.RandomString(5)
	tag, err := testQueries.CreateTag(context.Background(), ranTag)
	require.NoError(t, err)
	require.NotEmpty(t, tag)

	require.Equal(t, tag.TagName, ranTag)

	return tag
}

func TestCreateTag(t *testing.T) {
	CreateRandomTag(t)
}

func TestGetAllTags(t *testing.T) {
	var tag database.Tag
	for i := 0; i < 20; i++ {
		tag = CreateRandomTag(t)
	}

	tags, err := testQueries.GetAllTags(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, tags)

	require.Equal(t, tag.TagName, tags[len(tags)-1].TagName)
}
