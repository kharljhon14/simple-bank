package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/kharljhon14/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createTestTransaction(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	args := CreateTransactionParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoneyAmount(),
	}

	transfer, err := testStore.CreateTransaction(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	return transfer
}

func deleteTestTransaction(t *testing.T, id int64) {
	err := testStore.DeleteTransaction(context.Background(), id)
	require.NoError(t, err)
}

func TestCreateTransaction(t *testing.T) {
	transations := createTestTransaction(t)
	deleteTestTransaction(t, transations.ID)
}

func TestGetTransaction(t *testing.T) {
	transfer := createTestTransaction(t)

	transfer2, err := testStore.GetTransaction(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)

	deleteTestTransaction(t, transfer2.ID)
}

func TestListTransactions(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTestTransaction(t)
	}

	args := ListTransactionsParams{
		Limit:  5,
		Offset: 5,
	}

	transactions, err := testStore.ListTransactions(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}

	args = ListTransactionsParams{
		Limit:  10,
		Offset: 0,
	}

	transactions, _ = testStore.ListTransactions(context.Background(), args)
	for _, transaction := range transactions {
		deleteTestTransaction(t, transaction.ID)
	}
}

func TestDeleteTransaction(t *testing.T) {
	transaction := createTestTransaction(t)

	deleteTestTransaction(t, transaction.ID)

	transaction2, err := testStore.GetTransaction(context.Background(), transaction.ID)
	require.Error(t, err)
	errorMessage := fmt.Errorf("sql: %s", err)
	require.EqualError(t, errorMessage, sql.ErrNoRows.Error())
	require.Empty(t, transaction2)
}
