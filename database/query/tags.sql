-- name: CreateTag :one
INSERT INTO "Tags" (
  tag_name
) VALUES (
  $1
)
RETURNING *;

-- name: GetAllTags :many
select * from "Tags";

