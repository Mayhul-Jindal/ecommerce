-- name: CreateVerifyEmail :one
INSERT INTO "Verify_Emails" (
    user_id,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;


-- todo test this timeout feature how is this working
-- name: UpdateVerifyEmail :one
UPDATE "Verify_Emails"
SET
    is_used = TRUE
WHERE
    id = @id
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;


-- name: DeleteVerifyEmail :exec
delete from "Verify_Emails"
where user_id = $1;