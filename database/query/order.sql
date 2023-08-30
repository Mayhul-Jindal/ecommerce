-- name: AddOrder :one
INSERT INTO "Orders" (
  razorpay_order_id, user_id, total_money, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: GetOrderId :one
select id from "Orders"
where id = $1;



-- TODO use the expired_at > now() to handle some stuff