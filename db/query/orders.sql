-- name: CreateOrder :one
INSERT INTO "ordersDetails" (
   transfer_id, product_id, amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListOrders :many
SELECT * FROM "ordersDetails"
WHERE transfer_id = $1;

-- name: DeleteOrder :exec
DELETE FROM "ordersDetails"
WHERE order_id = $1;
