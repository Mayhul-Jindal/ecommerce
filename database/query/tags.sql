-- name: CreateTag :one
INSERT INTO "Tags" (
  id, tag_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllTags :many
select * from "Tags";


-- name: UpdateTag :one
update "Tags" set 
tag_name = $2 where
id = $1
returning *;

-- name: DeleteTag :exec
DELETE FROM "Tags"
WHERE id = $1;