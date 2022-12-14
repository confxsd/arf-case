-- name: CreateWallet :one
INSERT INTO wallets (
  user_id,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetWallet :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1;

-- name: GetWalletForUpdate :one
SELECT * FROM wallets
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetWalletByUserIdAndCurrency :one
SELECT * FROM wallets
WHERE user_id = $1 and currency = $2 LIMIT 1;

-- name: ListWallets :many
SELECT * FROM wallets
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateWallet :one
UPDATE wallets
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddWalletBalance :one
UPDATE wallets
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1;