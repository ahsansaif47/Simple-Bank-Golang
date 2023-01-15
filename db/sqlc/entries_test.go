package db

import (
	"context"
	"simple_bank_project/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {
	args := CreateEntryParams{
		AccountID: 1,
		Amount:    utils.RandomMoney(),
		CreatedAt: time.Time{},
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	// cheking no error running the transaction..
	require.NoError(t, err)
	// checking values equality..
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.WithinDuration(t, args.CreatedAt, entry.CreatedAt, time.Second)
	return entry
}
func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	NewEntry := CreateRandomEntry(t)
	RecEntry, err := testQueries.GetEntry(context.Background(), NewEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, RecEntry)
	require.Equal(t, RecEntry.AccountID, NewEntry.AccountID)
	require.Equal(t, RecEntry.Amount, NewEntry.Amount)
	require.WithinDuration(t, RecEntry.CreatedAt, NewEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomEntry(t)
	}
	args := ListEntriesParams{
		Limit:  3,
		Offset: 3,
	}
	EntryList, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		require.NotEmpty(t, EntryList)
	}
}
