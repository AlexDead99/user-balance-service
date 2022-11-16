-- name: CreateAccount :one
INSERT INTO accounts (
   owner, balance 
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateAccount :one
UPDATE accounts
  set balance = $2
WHERE account_id = $1
RETURNING *;
