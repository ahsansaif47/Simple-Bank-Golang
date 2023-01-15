package db

import (
	"context"
	"simple_bank_project/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	// checks if error is nill else it will fail the test..
	require.NoError(t, err)
	// cheking the returned object is not empty..
	require.NotEmpty(t, account)

	// verifying data..
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	// check if auto-fill columns are filled..
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	// check if there is no error..
	require.NoError(t, err)
	// check if the account object is not empty..
	require.NotEmpty(t, account2)

	// checking the fields of created account vs the account queried..
	require.Equal(t, account2.ID, account1.ID)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Currency, account1.Currency)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.CreatedAt, account1.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: utils.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Balance, args.Balance)
	require.Equal(t, account2.CreatedAt, account.CreatedAt)
	require.Equal(t, account2.Currency, account.Currency)
	require.Equal(t, account2.ID, account.ID)
	require.Equal(t, account2.Owner, account.Owner)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	del_account, err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), del_account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
	require.Equal(t, del_account.ID, account.ID)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomAccount(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	// checking if the accounts are not-empty
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
