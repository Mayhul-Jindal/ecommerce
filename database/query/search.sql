-- TODO fuzzy searching add karni hain isme fkin


-- name: SearchBooksV1 :many
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
select *,  
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
offset $3;


-- name: SearchBookV2 :many
