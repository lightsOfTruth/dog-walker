package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockdb "github.com/lightsOfTruth/dog-walker/db/mock"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
	"github.com/lightsOfTruth/dog-walker/helpers"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg           db.CreateUserParams
	plainPassword string
}

// Custom gomock matcher
func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	xArg, ok := x.(db.CreateUserParams)

	if !ok {
		return false
	}

	err := helpers.CheckPassword(e.plainPassword, xArg.Password)
	if err != nil {
		return false
	}

	eArgcreateUseParams := db.CreateUserParams{
		ID:        xArg.ID,
		FullName:  e.arg.FullName,
		Email:     e.arg.Email,
		Password:  xArg.Password,
		Contact:   e.arg.Contact,
		Dog:       e.arg.Dog,
		Address:   e.arg.Address,
		City:      e.arg.City,
		PostCode:  e.arg.PostCode,
		Longitude: e.arg.Longitude,
		Latitude:  e.arg.Latitude,
	}

	// In case, some value is nil
	if x == nil {
		return false
	}

	// Check if types assignable and convert them to common type
	x1Val := reflect.ValueOf(eArgcreateUseParams)
	x2Val := reflect.ValueOf(xArg)

	if x1Val.Type().AssignableTo(x2Val.Type()) {
		x1ValConverted := x1Val.Convert(x2Val.Type())
		return reflect.DeepEqual(x1ValConverted.Interface(), x2Val.Interface())
	}

	return false
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", e.arg, e.arg)
}

func EqCreateUserParams(arg db.CreateUserParams, plainPassword string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, plainPassword}

}

func TestCreateUser(t *testing.T) {

	userParams := db.CreateUserParams{
		ID:        uuid.New(),
		FullName:  "test user",
		Email:     "testemail@email.com",
		Password:  "testpassword",
		Contact:   "01111111111",
		Dog:       sql.NullInt32{Int32: 1, Valid: true},
		Address:   "50 Gotham",
		City:      "City",
		PostCode:  "D1 1AA",
		Longitude: "1268327832",
		Latitude:  "473987493",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(userParams, "testpassword")).
		Times(1).
		Return(db.User{
			ID:        userParams.ID,
			FullName:  "test user",
			Email:     "testemail@email.com",
			Contact:   "01111111111",
			Dog:       sql.NullInt32{Int32: 1, Valid: true},
			Address:   "50 Gotham",
			City:      "City",
			PostCode:  "D1 1AA",
			Longitude: "1268327832",
			Latitude:  "473987493",
		}, nil)

	server := newTestServer(t, store)
	recorder := httptest.NewRecorder()

	url := "/createuser"

	userParamsMarshalled := CreateUserRequestParams{
		FullName:  userParams.FullName,
		Email:     userParams.Email,
		Password:  userParams.Password,
		Contact:   userParams.Contact,
		Dog:       helpers.JsonNullInt32{NullInt32: userParams.Dog},
		Address:   userParams.Address,
		City:      userParams.City,
		PostCode:  userParams.PostCode,
		Longitude: userParams.Longitude,
		Latitude:  userParams.Latitude,
	}

	userParamsBytes, err := json.Marshal(userParamsMarshalled)

	if err != nil {
		println(err)

	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(userParamsBytes))
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	bodyData, bodyerr := ioutil.ReadAll(recorder.Body)

	if bodyerr != nil {
		println(bodyerr)
	}

	var returnedUser UserResponse
	err = json.Unmarshal(bodyData, &returnedUser)

	require.NoError(t, err)
	require.Equal(t, userParams.FullName, returnedUser.FullName)
	require.Equal(t, userParams.Dog.Int32, returnedUser.Dog)
}

// TODO: Write test for loginUser api
// func TestLoginUser(t *testing.T) {

// }
