-- name: CreateBook :one
INSERT INTO "Books" (
  title, author, tags_array, price, description
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetBookById :one
-- select b.id, b.title, b.author, b.price, array_agg(r.comment) as comments from "Books" b
-- join "Reviews" r on r.book_id = b.id
-- where b.id = $1
-- group by b.id, b.title, b.author, b.price;


-- name: UpdateBookDesc :one
UPDATE "Books"
set "description" = $2
WHERE "id" = $1
RETURNING *;

-- TODO: What happens when a book is deleted ?
-- name: DeleteBook :exec
DELETE FROM "Books"
WHERE id = $1;
