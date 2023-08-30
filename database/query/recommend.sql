-- name: GetHotSellingBooks: many
with cte as 
(
    select book_id from "Purchases"
    order by created_at desc
    group by book_id
    limit $1
    offset $2
)
select * from "Books" b
JOIN cte ON b.id = cte.book_id;



-- name: GetUserRecommendation: one
with cte as (
    select user_id, book_id from "Purchases" p
    where p.user_id = $1 and p.order_id = $2
)
select cte.user_id, string_agg(t.tag_name, ' ') as all_tags from "Books" b
join cte on cte.book_id = b.id
left join "Tags" t on t.id = ANY(b.tag_array)
group by cte.user_id;



-- WITH cte AS (
--     SELECT tags_array
--     FROM "Books"
--     WHERE id IN (1, 2, 3)
-- )
-- SELECT string_agg(t.tag_name, ' ') AS tags_array
-- FROM cte
-- LEFT JOIN "Tags" t ON t.id = ANY(cte.tags_array)
-- GROUP BY cte.tags_array;