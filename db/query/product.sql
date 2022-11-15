-- name: GetProduct :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE product_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateProduct :one
UPDATE products
  set amount = $2
WHERE product_id = $1
RETURNING *;
