-- name: CreateBook :one
INSERT INTO "Books" (
  title, author, tags, price, quantity, description
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetBook :one
select * from "Books"
where id = $1 limit 1;

-- name: GetBooks :many
select * from "Books"
where author = $1
limit $2
offset $3;

-- name: UpdateBookDesc :one
UPDATE "Books"
set "description" = $2
WHERE "id" = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM "Books"
WHERE id = $1;
