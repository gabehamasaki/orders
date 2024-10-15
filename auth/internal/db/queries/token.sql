-- name: InsertToken :one
INSERT INTO
    user_token (user_id, token, expires_at)
VALUES ($1, $2, $3) RETURNING id;
-- name: FindTokenByUserId :one
SELECT * FROM user_token WHERE user_id = $1;

-- name: DeletedTokenById :exec
DELETE FROM user_token WHERE id = $1;

-- name: DeletedTokenByUserId :exec
DELETE FROM user_token WHERE user_id = $1;

-- name: FindTokenByToken :one
SELECT * FROM user_token WHERE token = $1;

-- name: GetTokenIsExpiredByUserId :one
SELECT *
FROM user_token
WHERE
    user_id = $1
    AND expires_at < now();
