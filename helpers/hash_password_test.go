package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := randstr.String(20)
	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)

	require.NoError(t, err)

	wrongPassword := randstr.String(8)

	err = CheckPassword(wrongPassword, hashedPassword)

	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
