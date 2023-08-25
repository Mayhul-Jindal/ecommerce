// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: search.sql

package sqlc

import (
	"context"
)

const getTagsByBookId = `-- name: GetTagsByBookId :many
select tag_name from "Tags" as T
where T.id IN (
    select unnest(tags_array) from "Books" as B
    where B.id = $1
)
`

func (q *Queries) GetTagsByBookId(ctx context.Context, id int64) ([]string, error) {
	rows, err := q.db.Query(ctx, getTagsByBookId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var tag_name string
		if err := rows.Scan(&tag_name); err != nil {
			return nil, err
		}
		items = append(items, tag_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
