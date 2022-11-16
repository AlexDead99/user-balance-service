-- name: CreateTransfer :one
INSERT INTO transfers (
   user_id, service_id, total_price, description, status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE transfer_id = $1 LIMIT 1;

-- name: UpdateTransfer :one
UPDATE transfers
  set status = $2
WHERE transfer_id = $1
RETURNING *;

-- name: TransactionHistoryForUser :many
SELECT * FROM transfers
WHERE user_id = $1 AND EXTRACT(MONTH FROM created_at) >= $1 AND EXTRACT(YEAR FROM created_at) >= $2 AND status NOT IN ('Success')

-- name: GeneralReport :many
SELECT * FROM transfers
WHERE EXTRACT(MONTH FROM created_at) = $1 AND EXTRACT(YEAR FROM created_at) = $2 AND status NOT IN ('Success')

