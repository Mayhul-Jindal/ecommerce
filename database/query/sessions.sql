-- name: CreateSession :one
INSERT INTO "Sessions" (
  id,
  user_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM "Sessions"
WHERE id = $1 LIMIT 1;

-- name: DeleteSession :exec
delete from "Sessions"
where user_id = $1;