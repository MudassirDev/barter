-- name: CreateUser :one
INSERT INTO users (
  id, username, first_name, last_name, email, password, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetUserWithEmail :one
SELECT * FROM users WHERE email = ?;
