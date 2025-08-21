-- name: CreateUser :one
INSERT INTO users (
  id, username, first_name, last_name, email, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;
