package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/kharljhon14/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createTestEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoneyAmount(),
	}

	entry, err := testStore.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	return entry

}

func deleteTestEntry(t *testing.T, id int64) {
	err := testStore.DeleteEntry(context.Background(), id)
	require.NoError(t, err)
}

func TestCreateEntry(t *testing.T) {
	entry := createTestEntry(t)

	deleteTestEntry(t, entry.ID)
}

func TestGetEntry(t *testing.T) {
	entry := createTestEntry(t)

	entry2, err := testStore.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)

	deleteTestEntry(t, entry2.ID)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestEntry(t)
	}

	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testStore.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

	args = ListEntriesParams{
		Limit:  10,
		Offset: 0,
	}

	entries, _ = testStore.ListEntries(context.Background(), args)
	for _, entry := range entries {
		deleteTestEntry(t, entry.ID)
	}
}

func TestDeleteEntry(t *testing.T) {
	entry := createTestEntry(t)

	deleteTestEntry(t, entry.ID)

	entry2, err := testStore.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	errorMessage := fmt.Errorf("sql: %s", err)
	require.EqualError(t, errorMessage, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
