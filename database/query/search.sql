-- name: SearchBooksV2 :many
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
  SELECT *,
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
SELECT *,
  GREATEST(title_difference, author_difference, description_difference, tags_difference) AS max_difference
FROM final_cte
ORDER BY ranks DESC, max_difference DESC, tags_difference DESC, title_difference DESC, author_difference DESC, description_difference DESC
LIMIT $3
OFFSET $4;

