
-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    balance,
    currency
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE accounts 
SET owner = $2,balance = $3,currency = $4
WHERE id = $1
RETURNING *;

-- name: ListAccount :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;
