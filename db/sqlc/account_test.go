package db

import (
	"context"
	"simple_bank_project/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
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
}
