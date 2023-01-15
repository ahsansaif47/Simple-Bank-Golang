package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	args := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   3,
		Amount:        2000,
		CreatedAt:     time.Time{},
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, args.Amount, transfer.Amount)
	require.WithinDuration(t, args.CreatedAt, transfer.CreatedAt, time.Second)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	RetTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, RetTransfer)
	require.Equal(t, transfer.Amount, RetTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, RetTransfer.CreatedAt, time.Second)
	require.Equal(t, transfer.FromAccountID, RetTransfer.FromAccountID)
	require.Equal(t, transfer.ID, RetTransfer.ID)
	require.Equal(t, transfer.ToAccountID, RetTransfer.ToAccountID)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomTransfer(t)
	}
	args := ListTransfersParams{
		Limit:  3,
		Offset: 3,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		require.NotEmpty(t, transfers)
	}
}
