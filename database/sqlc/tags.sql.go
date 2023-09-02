// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: tags.sql

package database

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO "Tags" (
  id, tag_name
) VALUES (
  $1, $2
)
RETURNING id, tag_name
`

type CreateTagParams struct {
	ID      int32  `json:"id"`
	TagName string `json:"tag_name"`
}

func (q *Queries) CreateTag(ctx context.Context, arg CreateTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, createTag, arg.ID, arg.TagName)
	var i Tag
	err := row.Scan(&i.ID, &i.TagName)
	return i, err
}

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM "Tags"
WHERE id = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTag, id)
	return err
}

const getAllTags = `-- name: GetAllTags :many
select id, tag_name from "Tags"
`

func (q *Queries) GetAllTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.Query(ctx, getAllTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.TagName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTag = `-- name: UpdateTag :one
update "Tags" set 
tag_name = $2 where
id = $1
returning id, tag_name
`

type UpdateTagParams struct {
	ID      int32  `json:"id"`
	TagName string `json:"tag_name"`
}

func (q *Queries) UpdateTag(ctx context.Context, arg UpdateTagParams) (Tag, error) {
	row := q.db.QueryRow(ctx, updateTag, arg.ID, arg.TagName)
	var i Tag
	err := row.Scan(&i.ID, &i.TagName)
	return i, err
}
