-- name: CreateUser :exec
INSERT INTO users (id, username, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- name: GetUserByEmail :one
SELECT id, username, email, password, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id, username, email, password, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2, email = $3, password = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;