-- name: CreateTag :one
INSERT INTO "Tags" (
  id, tag_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllTags :many
select * from "Tags";

