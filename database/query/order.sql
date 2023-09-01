-- name: AddOrder :one
INSERT INTO "Orders" (
  razorpay_order_id, user_id, total_money, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetOrderById :one
select * from "Orders"
where id = $1 and user_id = $2;

-- name: UpdateOrder :one
UPDATE "Orders"
set "status" = $1
WHERE "id" = $2 and user_id = $3
RETURNING *;

-- name: DeleteOrder :exec
delete from "Orders"
where user_id = $1;



-- TODO use the expired_at > now() to handle some stuff