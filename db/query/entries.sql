-- name: CreateEntry :one
INSERT INTO entries(
        account_id,
        amount,
        created_at
    )
VALUES ($1, $2, $3, $4)
Returning *;
-- name: GetEntry :one
SELECT *
FROM entries
WHERE id = $1
LIMIT 1;
-- name: ListEntries :many
SELECT *
FROM entries
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;