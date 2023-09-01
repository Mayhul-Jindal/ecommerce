// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: search.sql

package database

import (
	"context"
)

const searchBooksV1 = `-- name: SearchBooksV1 :many


with book_search_cte as 
(
    SELECT
        b.id,
        b.title,
        b.author,
        array_agg(t.tag_name) AS tags_array,
        r.average_ratings as ratings,
        b.description
    FROM "Books" AS b
    LEFT JOIN "Tags" AS t ON t.id = ANY(b.tags_array)
    LEFT JOIN (
        select book_id, avg(rating) as average_ratings from "Reviews"
        group by book_id
        ) as r on r.book_id = b.id
    GROUP BY
        b.id, b.title, b.author, r.average_ratings, b.description
)
select id, title, author, tags_array, ratings, description,  
ts_rank(
	to_tsvector('english', title) || ' ' ||
	to_tsvector('english', author) || ' ' ||
	setweight(to_tsvector('english', array_to_string(COALESCE(tags_array, '{}'), ' ')), 'A') || ' ' ||
	to_tsvector('english', description),
	websearch_to_tsquery('english', $1)
  ) AS ranks
from book_search_cte
order by ranks desc
limit $2
offset $3
`

type SearchBooksV1Params struct {
	WebsearchToTsquery string `json:"websearch_to_tsquery"`
	Limit              int32  `json:"limit"`
	Offset             int32  `json:"offset"`
}

type SearchBooksV1Row struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Author      string      `json:"author"`
	TagsArray   interface{} `json:"tags_array"`
	Ratings     float64     `json:"ratings"`
	Description string      `json:"description"`
	Ranks       float32     `json:"ranks"`
}

// TODO fuzzy searching add karni hain isme fkin
func (q *Queries) SearchBooksV1(ctx context.Context, arg SearchBooksV1Params) ([]SearchBooksV1Row, error) {
	rows, err := q.db.Query(ctx, searchBooksV1, arg.WebsearchToTsquery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchBooksV1Row{}
	for rows.Next() {
		var i SearchBooksV1Row
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.TagsArray,
			&i.Ratings,
			&i.Description,
			&i.Ranks,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchBooksV2 = `-- name: SearchBooksV2 :many
WITH book_search_cte AS (
    SELECT
        b.id,
        b.title,
        b.author,
        array_agg(t.tag_name) AS tags_array,
        r.average_ratings AS ratings,
        b.description
    FROM "Books" AS b
    LEFT JOIN "Tags" AS t ON t.id = ANY(b.tags_array)
    LEFT JOIN (
        SELECT book_id, AVG(rating) AS average_ratings FROM "Reviews"
        GROUP BY book_id
    ) AS r ON r.book_id = b.id
    GROUP BY b.id, b.title, b.author, r.average_ratings, b.description
), 
final_cte AS (
  SELECT id, title, author, tags_array, ratings, description,
    ts_rank(
      to_tsvector('english', title) || ' ' ||
      to_tsvector('english', author) || ' ' ||
      setweight(to_tsvector('english', array_to_string(COALESCE(tags_array, '{}'), ' ')), 'A') || ' ' ||
      to_tsvector('english', description),
      websearch_to_tsquery('english', $1)
    ) AS ranks,
    difference(array_to_string(COALESCE(tags_array, '{}'), ' '), $2) AS tags_difference,
    difference(title, $2) AS title_difference,
    difference(author, $2) AS author_difference,
    difference(description, $2) AS description_difference
  FROM book_search_cte
)
SELECT id, title, author, tags_array, ratings, description, ranks, tags_difference, title_difference, author_difference, description_difference,
  GREATEST(title_difference, author_difference, description_difference, tags_difference) AS max_difference
FROM final_cte
ORDER BY ranks DESC, max_difference DESC, tags_difference DESC, title_difference DESC, author_difference DESC, description_difference DESC
LIMIT $3
OFFSET $4
`

type SearchBooksV2Params struct {
	WebsearchToTsquery string `json:"websearch_to_tsquery"`
	Difference         string `json:"difference"`
	Limit              int32  `json:"limit"`
	Offset             int32  `json:"offset"`
}

type SearchBooksV2Row struct {
	ID                    int64       `json:"id"`
	Title                 string      `json:"title"`
	Author                string      `json:"author"`
	TagsArray             interface{} `json:"tags_array"`
	Ratings               interface{}     `json:"ratings"`
	Description           string      `json:"description"`
	Ranks                 float32     `json:"ranks"`
	TagsDifference        int32       `json:"tags_difference"`
	TitleDifference       int32       `json:"title_difference"`
	AuthorDifference      int32       `json:"author_difference"`
	DescriptionDifference int32       `json:"description_difference"`
	MaxDifference         interface{} `json:"max_difference"`
}

func (q *Queries) SearchBooksV2(ctx context.Context, arg SearchBooksV2Params) ([]SearchBooksV2Row, error) {
	rows, err := q.db.Query(ctx, searchBooksV2,
		arg.WebsearchToTsquery,
		arg.Difference,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchBooksV2Row{}
	for rows.Next() {
		var i SearchBooksV2Row
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.TagsArray,
			&i.Ratings,
			&i.Description,
			&i.Ranks,
			&i.TagsDifference,
			&i.TitleDifference,
			&i.AuthorDifference,
			&i.DescriptionDifference,
			&i.MaxDifference,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
