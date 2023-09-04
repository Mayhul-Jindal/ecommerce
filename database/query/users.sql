-- name: CreateUser :one
INSERT INTO "Users" (
  username, email, hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
select * from "Users"
where id = $1 and username = $2 and not is_deleted
limit 1;

-- name: GetUserByUsername :one
select * from "Users"
where username = $1
limit 1;


-- name: GetUserById :one
select * from "Users"
where id = $1
limit 1;

-- name: UpdateUser :one
UPDATE "Users"
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified),
  is_active = COALESCE(sqlc.narg(is_active), is_active),
  deactivated_at = COALESCE(sqlc.narg(deactivated_at), deactivated_at),
  is_deleted = COALESCE(sqlc.narg(is_deleted), is_deleted),
  deleted_at = COALESCE(sqlc.narg(deleted_at), deleted_at)
WHERE
  id = sqlc.arg(id)
RETURNING *;


-- name: CheckAdmin :one
select * from "Users"
where id = $1 and is_admin;

-- name: CheckEmailVerified :one
select * from "Users"
where id = $1 and is_email_verified;


-- name: DeleteUser :exec
delete from "Users"
where id = $1;