-- name: GetService :one
SELECT * FROM services
WHERE service_id = $1 LIMIT 1;
