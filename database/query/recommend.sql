-- name: GetHotSellingBooks :many
SELECT b.title, b.author, b.price FROM "Books" b
JOIN "Purchases" p ON b.id = p.book_id
GROUP BY p.book_id, b.title, b.author, b.price
ORDER BY COUNT(*) DESC
LIMIT $1
OFFSET $2;


-- name: GetUserRecommendations :many
with cte as (
	SELECT string_agg(t.tag_name, ' ') as all_tags FROM "Purchases" p
	JOIN "Books" b ON p.book_id = b.id
	LEFT JOIN "Tags" t ON t.id = ANY(b.tags_array)
	WHERE p.user_id = $1 AND p.order_id = $2
	GROUP BY p.user_id, t.tag_name
)
select string_agg(all_tags, ' ') from cte;



