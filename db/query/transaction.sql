-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, transaction_type, amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY amount
LIMIT $1
OFFSET $2;