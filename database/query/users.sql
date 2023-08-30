-- name: CreateUser :one
INSERT INTO "Users" (
  username, email, hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
select * from "Users"
where id = $1 and username = $2
limit 1;

-- name: UpdateUser :one
UPDATE "Users"
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeactivateUser :one
UPDATE "Users" 
SET is_active = false 
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one 
UPDATE "Users" 
SET IsDeleted = true 
WHERE id = $1
RETURNING *;;



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
