-- name: CreatePurchase :one
INSERT INTO "Purchases" (
  user_id, book_id, order_id
) VALUES (
  $1, $2, $3
)
RETURNING *;


-- name: CheckBookPurchased :one
select * from "Purchases"
where user_id = $1 and book_id = $2;

-- todo this is to get bought books at a single place
-- name: GetPurchasedBooks :many
with cte as (
    select * from "Purchases"
    where user_id = $1
)
select * from "Books" b
join cte on cte.book_id = b.id
limit $2
offset $3;

-- name: DeletePurchase :exec
DELETE FROM "Purchases"
WHERE user_id = $1;


