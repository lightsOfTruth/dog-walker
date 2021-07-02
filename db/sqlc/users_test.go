package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var userInserted CreateUserParams

func TestCreateUser(t *testing.T) {
	userArgs := CreateUserParams{
		ID: uuid.New(),
		FullName: "Steve Austin",
		Contact: "01111111111",
		Dog: sql.NullInt32{Int32: 1, Valid: true},
		Address: "50 Gotham",
		City: "City",
		PostCode: "D1 1AA",
		Longitude: "1268327832",
		Latitude: "473987493",
	}

	userInserted = userArgs

	user, err := testQueries.CreateUser(context.Background(), userArgs)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, userArgs.FullName, user.FullName)
	require.Equal(t, userArgs.Contact, user.Contact)
	require.Equal(t, userArgs.Dog, user.Dog)
	require.Equal(t, userArgs.Address, user.Address)
	require.Equal(t, userArgs.City, user.City)
	require.Equal(t, userArgs.PostCode, user.PostCode)
	require.Equal(t, userArgs.Longitude, user.Longitude)
	require.Equal(t, userArgs.Latitude, user.Latitude)

}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), userInserted.ID)

	require.NoError(t, err)
	require.Equal(t, userInserted.ID, user.ID)
}

func TestUpdateUser(t *testing.T) {

	userInserted.City = "Metropolis"

	user, err := testQueries.UpdateUser(context.Background(),
	 UpdateUserParams{ID: userInserted.ID,
		 Contact: userInserted.Contact,
		  Address: userInserted.Address,
		   City: userInserted.City,
		   PostCode: userInserted.PostCode,
		   Dog: userInserted.Dog,
		})

	require.NoError(t, err)
	require.Equal(t, userInserted.City, user.City)
}


func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), userInserted.ID)

	require.NoError(t, err)


}
