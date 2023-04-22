package db

import (
	"context"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{AccountID: account.ID,
		Amount: util.RandomMoney()}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.Amount)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	randomEntry := createRandomEntry(t, account)
	entry, err := testQueries.GetEntry(context.Background(), randomEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, randomEntry.ID, entry.ID)
	require.Equal(t, randomEntry.AccountID, entry.AccountID)
	require.Equal(t, randomEntry.Amount, entry.Amount)
	require.WithinDuration(t, randomEntry.CreatedAt, entry.CreatedAt, time.Second)
}

func TestGetListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 6; i++ {
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams{AccountID: account.ID, Limit: 3, Offset: 3}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 3)
	for _, entry := range entries{
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
}
