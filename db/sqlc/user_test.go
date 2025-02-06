package db

import (
	"context"
	"testing"
	"time"

	"github.com/kharljhon14/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) User {
	hashedPassword, err := util.HashedPassword(util.RandomString(8))
	require.NoError(t, err)

	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.NotZero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUser(t *testing.T) {
	user := createTestUser(t)

	user1, err := testStore.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.HashedPassword, user1.HashedPassword)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.Email, user1.Email)
	require.WithinDuration(t, user.PasswordChangedAt.Time, user1.PasswordChangedAt.Time, time.Second)
	require.WithinDuration(t, user.CreatedAt.Time, user1.CreatedAt.Time, time.Second)
}
