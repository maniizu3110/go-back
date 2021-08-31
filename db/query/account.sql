
-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: CreateAccount :one
INSERT INTO account (
    owner,
    balance,
    currency
) VALUES (
  $1, $2, $3
) RETURNING *;


-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE account SET balance = $2
WHERE id = $1
RETURNING *;

-- name: ListAccount :many
SELECT * FROM account
ORDER BY id
LIMIT $1
OFFSET $2;