-- name: CreateBook :one
INSERT INTO "Books" (
  title, author, tags_array, price, description, download_link
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetBookById :one
select title, author, price, description, download_link from "Books"
where id = $1;

-- name: UpdateBook :one
UPDATE "Books"
SET
  title = COALESCE(sqlc.narg(title), title),
  author = COALESCE(sqlc.narg(author), author),
  tags_array = COALESCE(sqlc.narg(tags_array), tags_array),
  price = COALESCE(sqlc.narg(price), price),
  description = COALESCE(sqlc.narg(description), description),
  download_link = COALESCE(sqlc.narg(download_link), download_link)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM "Books"
WHERE id = $1;
