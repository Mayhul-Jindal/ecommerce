-- name: CreateReview :one
INSERT INTO "Reviews" (
  user_id, book_id, rating, comment
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetReviewsByBookId :many
select * from "Reviews"
where book_id = $1
limit $2
offset $3;

-- name: UpdateReview :one
UPDATE "Reviews"
set "rating" = $2, "comment" = $3
WHERE "id" = $1
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM "Reviews"
WHERE id = $1;


