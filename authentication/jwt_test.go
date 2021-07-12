package authentication

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func TestJWT(t *testing.T) {
	jwtCreator, err := NewJWTCreator(randstr.String(32))
	require.NoError(t, err)

	token, err := jwtCreator.CreateToken("test user", time.Minute)
	require.NoError(t, err)

	payload, err := jwtCreator.VerifyToken(token)
	require.NoError(t, err)

	require.NotEmpty(t, payload)
	require.Equal(t, payload.Username, "test user")

}

func TestJWTExpiredToken(t *testing.T) {
	jwtCreator, err := NewJWTCreator(randstr.String(32))
	require.NoError(t, err)

	token, err := jwtCreator.CreateToken("test user2", time.Second)
	require.NoError(t, err)

	time.Sleep(time.Duration(time.Second * 2))

	payload, err := jwtCreator.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrorExpiredToken.Error())
	require.Nil(t, payload)
}
