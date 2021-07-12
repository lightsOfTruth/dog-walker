package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/lightsOfTruth/dog-walker/helpers"
	"github.com/stretchr/testify/require"
)

var userInserted CreateUserParams

const TEST_PASSWORD = "testpassword"
const TEST_EMAIL = "testemail@email.com"

func TestCreateUser(t *testing.T) {

	hashedPassword, err := helpers.HashPassword(TEST_PASSWORD)
	require.NoError(t, err)

	userInserted = CreateUserParams{
		ID:        uuid.New(),
		FullName:  "test user",
		Email:     TEST_EMAIL,
		Password:  hashedPassword,
		Contact:   "01111111111",
		Dog:       sql.NullInt32{Int32: 1, Valid: true},
		Address:   "50 Gotham",
		City:      "City",
		PostCode:  "D1 1AA",
		Longitude: "1268327832",
		Latitude:  "473987493",
	}

	fmt.Printf("userinsertID is %v", userInserted.ID)
	user, err := testQueries.CreateUser(context.Background(), userInserted)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, userInserted.FullName, user.FullName)
	require.Equal(t, userInserted.Contact, user.Contact)
	require.Equal(t, userInserted.Dog, user.Dog)
	require.Equal(t, userInserted.Address, user.Address)
	require.Equal(t, userInserted.City, user.City)
	require.Equal(t, userInserted.PostCode, user.PostCode)
	require.Equal(t, userInserted.Longitude, user.Longitude)
	require.Equal(t, userInserted.Latitude, user.Latitude)

}

func TestGetUserByID(t *testing.T) {
	fmt.Printf("userinsertID in get is %v\n", userInserted.ID)
	// parsed, _ := uuid.Parse("66d7d7da-f078-4378-852e-680183d21f63")
	// println(parsed.String())
	user, err := testQueries.GetUserById(context.Background(), userInserted.ID)

	require.NoError(t, err)
	require.Equal(t, userInserted.ID, user.ID)
}
func TestGetUserByEmail(t *testing.T) {

	user, err := testQueries.GetUserByEmail(context.Background(), TEST_EMAIL)

	require.NoError(t, err)
	require.Equal(t, user.Email, TEST_EMAIL)
}

func TestUpdateUser(t *testing.T) {

	userInserted.City = "Metropolis"

	user, err := testQueries.UpdateUser(context.Background(),
		UpdateUserParams{ID: userInserted.ID,
			Contact:  userInserted.Contact,
			Address:  userInserted.Address,
			City:     userInserted.City,
			PostCode: userInserted.PostCode,
			Dog:      userInserted.Dog,
		})

	require.NoError(t, err)
	require.Equal(t, userInserted.City, user.City)
}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), userInserted.ID)

	require.NoError(t, err)

}
