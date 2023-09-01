-- name: GetHotSellingBooks :many
with cte as 
(
    select book_id from "Purchases"
    group by book_id
    order by created_at desc
    limit $1
    offset $2
)
select title, author, price from "Books" b
JOIN cte ON b.id = cte.book_id;

-- name: GetUserRecommendations :many
with cte as 
(
    select user_id, book_id from "Purchases" p
    where p.user_id = $1 and p.order_id = $2
)
select cte.user_id, string_agg(t.tag_name, ' ') from "Books" b
join cte on cte.book_id = b.id
left join "Tags" t on t.id = ANY(b.tag_array)
group by cte.user_id, t.tag_name
limit $3
offset $4;



