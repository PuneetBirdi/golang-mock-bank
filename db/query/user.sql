-- name: CreateUser :one
INSERT INTO users (
    full_name,
    hashed_password,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE 
	id = $1 OR 
	email = $2
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET full_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE ID $1;
