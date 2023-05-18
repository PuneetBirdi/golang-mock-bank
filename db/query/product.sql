-- name: GetProduct :one
SELECT * FROM products
WHERE ID = $1 LIMIT 1;
