// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: wallet.sql

package arfcasesqlc

import (
	"context"
)

const addWalletBalance = `-- name: AddWalletBalance :one
UPDATE wallets
SET balance = balance + $1
WHERE id = $2
RETURNING id, user_id, balance, currency, created_at
`

type AddWalletBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) AddWalletBalance(ctx context.Context, arg AddWalletBalanceParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, addWalletBalance, arg.Amount, arg.ID)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const createWallet = `-- name: CreateWallet :one
INSERT INTO wallets (
  user_id,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING id, user_id, balance, currency, created_at
`

type CreateWalletParams struct {
	UserID   int64  `json:"user_id"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, createWallet, arg.UserID, arg.Balance, arg.Currency)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteWallet = `-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1
`

func (q *Queries) DeleteWallet(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteWallet, id)
	return err
}

const getWallet = `-- name: GetWallet :one
SELECT id, user_id, balance, currency, created_at FROM wallets
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWallet(ctx context.Context, id int64) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, getWallet, id)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getWalletForUpdate = `-- name: GetWalletForUpdate :one
SELECT id, user_id, balance, currency, created_at FROM wallets
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetWalletForUpdate(ctx context.Context, id int64) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, getWalletForUpdate, id)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listWallets = `-- name: ListWallets :many
SELECT id, user_id, balance, currency, created_at FROM wallets
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListWalletsParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListWallets(ctx context.Context, arg ListWalletsParams) ([]Wallet, error) {
	rows, err := q.db.QueryContext(ctx, listWallets, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wallet{}
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateWallet = `-- name: UpdateWallet :one
UPDATE wallets
SET balance = $2
WHERE id = $1
RETURNING id, user_id, balance, currency, created_at
`

type UpdateWalletParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (q *Queries) UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, updateWallet, arg.ID, arg.Balance)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}
