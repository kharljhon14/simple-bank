package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/kharljhon14/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoneyAmount(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueires.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func deleteTestAccount(t *testing.T, id int64) {
	err := testQueires.DeleteAccount(context.Background(), id)
	require.NoError(t, err)
}

func TestCreateAccount(t *testing.T) {
	account := createRandomAccount(t)

	deleteTestAccount(t, account.ID)
}

func TestGetAccount(t *testing.T) {
	// Create Account
	account := createRandomAccount(t)

	// Get account
	account2, err := testQueires.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.WithinDuration(
		t,
		account.CreatedAt.Time,
		account2.CreatedAt.Time,
		time.Second,
	)

	deleteTestAccount(t, account2.ID)
}

func TestUpdateAccount(t *testing.T) {
	// Create account
	account := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoneyAmount(),
	}

	account2, err := testQueires.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.WithinDuration(
		t,
		account.CreatedAt.Time,
		account2.CreatedAt.Time,
		time.Second,
	)

	deleteTestAccount(t, account2.ID)

}

func TestDeleteAccount(t *testing.T) {
	// Create account
	account := createRandomAccount(t)
	deleteTestAccount(t, account.ID)

	account2, err := testQueires.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	errorMessage := fmt.Errorf("sql: %s", err)
	require.EqualError(t, errorMessage, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueires.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

	args = ListAccountsParams{
		Limit:  10,
		Offset: 0,
	}
	accounts, _ = testQueires.ListAccounts(context.Background(), args)
	for _, account := range accounts {
		deleteTestAccount(t, account.ID)
	}
}
