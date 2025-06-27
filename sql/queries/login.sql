-- name: CheckEmail :one
SELECT * FROM users
WHERE email = $1;