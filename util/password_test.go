package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)

	hashed_password, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password)

	err = CheckPassword(password, hashed_password)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashed_password)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
