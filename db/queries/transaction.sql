-- name: CreateTransaction :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount   
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transfers
WHERE id = $1;

-- name: ListTransactions :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteTransaction :exec
DELETE FROM transfers
WHERE id = $1;