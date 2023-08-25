-- name: CreateUser :one
INSERT INTO "Users" (
  username, email, hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
select * from "Users"
where id = $1 limit 1;

-- -- name: GetUsers :many
-- select * from "Users"
-- limit $1
-- offset $2;

-- -- name: UpdateBookDesc :one
-- UPDATE "Users"
-- set "description" = $2
-- WHERE "id" = $1
-- RETURNING *;

-- -- name: DeleteBook :exec
-- DELETE FROM "Users"
-- WHERE id = $1;
