package db

import (
	"context"
	"database/sql"
	"fmt"
)

// this struct provides all the functions to execute database queries and transactions..
type Store struct {
	*Queries
	db *sql.DB
}

// Create a new store..
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Contains necessary parameters to transfer from 1 account to other..
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// func to execute generate database transaction..
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// TxOptions allows us to set isolation levels from a transaction..
	// if not set then the default isolation level is used (i.e. 0)
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// getting queries object using New() function
	q := New(tx)
	err = fn(q)
	// checking if there is an error executing callback function..
	if err != nil {
		// checking for rollback error..
		// if there is rollback error return rollback error and begintx error..
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		// if there is function execution error return that..
		return err
	}
	// else commit the transaction.. which returns an error.. (nil/error)
	return tx.Commit()
}

// Transfer performs money transfer from one acc. to another..
// Creates a transfer record, entry and account record. Updates the balance within a single transaction..
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx,
			CreateTransferParams{
				FromAccountID: arg.FromAccountID,
				ToAccountID:   arg.ToAccountID,
				Amount:        arg.Amount,
			})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// Update account balance
		// Updating account balance in an order(to avoid deadlocks)..
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return result, err
	}
	return result, nil
}

func addMoney(
	ctx context.Context,
	q *Queries,
	fromAccountID int64,
	amount1 int64,
	toAccountID int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount1,
		ID:     fromAccountID,
	})
	if err != nil {
		return
	}
	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount2,
		ID:     toAccountID,
	})
	return
}
