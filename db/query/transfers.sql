-- name: CreateEntry :one
INSERT INTO transfers(
        from_account_id,
        to_account_id,
        amount,
        created_at
    )
VALUES ($1, $2, $3, $4)
Returning *;
-- name: GetEntry :one
SELECT *
FROM transfers
WHERE id = $1
LIMIT 1;
-- name: ListEntries :many
SELECT *
FROM transfers
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: DeleteEntry :exec
DELETE FROM transfers
WHERE id = $1;