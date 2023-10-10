-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY first_name;

-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, password, email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
  set first_name = $2,
  last_name = $3,
  password = $4,
  email = $5
WHERE id = $1
RETURNING *;