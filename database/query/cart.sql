-- name: AddToCart :one
INSERT INTO "Carts" (
  user_id, book_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetCartItemsByUserId :many
select b.id, b.title, b.price from "Carts" c
join "Books" b on b.id = c.book_id 
where c.user_id = $1;

-- name: GetTotalCartAmountById :one
select sum(b.price) as total_money from "Carts" c
join "Books" b on c.book_id = b.id
where c.user_id = $1
group by user_id;

-- name: DeleteCartItem :exec
DELETE FROM "Carts"
WHERE user_id = $1 and book_id = $2;


-- name: DeleteCartOfUser :exec
DELETE FROM "Carts"
WHERE user_id = $1;


