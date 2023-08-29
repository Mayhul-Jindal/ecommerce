-- name: AddOrder :one
INSERT INTO "Orders" (
  razorpay_order_id, user_id, total_money, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- nmae: AddOrderItems :many
INSERT INTO "Order_Lines" (
  book_id, order_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetOrderId :one
select id from "Orders"
where id = $1;