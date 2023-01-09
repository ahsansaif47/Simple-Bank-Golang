-- name: CreateAccount :one
INSERT INTO accounts(
        owner,
        balance,
        currency,
        created_at
    )
VALUES ($1, $2, $3, $4)
Returning *;
-- name: GetAccount :one
SELECT *
FROM accounts
WHERE id = $1
LIMIT 1;
-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: UpdateAccount :exec
UPDATE accounts
SET balance = $2
WHERE id = $1;
-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;