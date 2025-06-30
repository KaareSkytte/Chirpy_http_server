-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetUserFromRefreshToken :one
SELECT users.id, users.created_at, users.updated_at, users.email, users.hashed_password,
       refresh_tokens.expires_at, refresh_tokens.revoked_at
FROM users
INNER JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens 
SET revoked_at = $2, updated_at = $3
WHERE token = $1;

