-- name: CreateBook :one
INSERT INTO "Books" (
  title, author, tags_array, price, quantity, description
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- TODO rating with reviews bhi merge hone chahiye isme ideally
-- name: GetBook :one
select * from "Books"
where id = $1 limit 1;

-- name: GetBooks :many
select * from "Books"
limit $1
offset $2;

-- name: UpdateBookDesc :one
UPDATE "Books"
set "description" = $2
WHERE "id" = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM "Books"
WHERE id = $1;
