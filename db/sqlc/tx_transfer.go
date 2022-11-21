package arfcasesqlc

import "context"

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromWalletID int64   `json:"from_wallet_id"`
	ToWalletID   int64   `json:"to_wallet_id"`
	Amount       float64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer1 Transfer `json:"transfer1"`
	Transfer2 Transfer `json:"transfer2"`
}

// TransferTx performs a money transfer from one wallet to the other.
func (store *SQLStore) TransferTx(ctx context.Context, arg [2]TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer1, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromWalletID: arg[0].FromWalletID,
			ToWalletID:   arg[0].ToWalletID,
			Amount:       arg[0].Amount,
		})
		if err != nil {
			return err
		}

		result.Transfer2, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromWalletID: arg[1].FromWalletID,
			ToWalletID:   arg[1].ToWalletID,
			Amount:       arg[1].Amount,
		})
		if err != nil {
			return err
		}

		_, _, err = addMoney(ctx, q, arg[0].FromWalletID, -arg[0].Amount, arg[0].ToWalletID, arg[0].Amount)
		if err != nil {
			return err
		}
		_, _, err = addMoney(ctx, q, arg[1].FromWalletID, -arg[1].Amount, arg[1].ToWalletID, arg[1].Amount)
		if err != nil {
			return err
		}
		return err
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	walletID1 int64,
	amount1 float64,
	walletID2 int64,
	amount2 float64,
) (wallet1 Wallet, wallet2 Wallet, err error) {
	wallet1, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
		ID:     walletID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	wallet2, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
		ID:     walletID2,
		Amount: amount2,
	})
	return
}
