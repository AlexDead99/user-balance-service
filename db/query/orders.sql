-- name: CreateOrder :one
INSERT INTO "ordersDetails" (
   transfer_id, product_id, amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ListOrders :many
SELECT * FROM "ordersDetails"
ORDER BY transfer_id;

