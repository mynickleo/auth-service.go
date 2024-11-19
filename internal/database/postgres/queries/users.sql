-- name: CreateUser :exec
INSERT INTO users (id, full_name, email, password)
VALUES ($1, $2, $3, $4);

-- name: GetUsers :many
SELECT id, email, full_name, created_at, updated_at
FROM users;

-- name: GetUserByID :one
SELECT id, email, full_name, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, password, full_name, created_at, updated_at
FROM users
WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users
SET email = $1, password = $2, full_name = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;