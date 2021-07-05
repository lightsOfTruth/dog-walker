package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mockdb "github.com/lightsOfTruth/dog-walker/db/mock"
	db "github.com/lightsOfTruth/dog-walker/db/sqlc"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

	userParams := struct {
		ID        uuid.UUID `json:"id"`
		FullName  string    `json:"full_name"`
		Contact   string    `json:"contact"`
		Dog       int       `json:"dog"`
		Address   string    `json:"address"`
		City      string    `json:"city"`
		PostCode  string    `json:"post_code"`
		Longitude string    `json:"longitude"`
		Latitude  string    `json:"latitude"`
	}{
		ID:        uuid.New(),
		FullName:  "test user",
		Contact:   "01111111111",
		Dog:       1,
		Address:   "50 Gotham",
		City:      "City",
		PostCode:  "D1 1AA",
		Longitude: "1268327832",
		Latitude:  "473987493",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
		Times(1).
		Return(db.User{
			ID:        userParams.ID,
			FullName:  "mock test user",
			Contact:   "01111111111",
			Dog:       sql.NullInt32{Int32: 1, Valid: true},
			Address:   "50 Gotham",
			City:      "City",
			PostCode:  "D1 1AA",
			Longitude: "1268327832",
			Latitude:  "473987493",
		}, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := "/createuser"

	b, err := json.Marshal(userParams)

	if err != nil {
		println(err)

	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	bodyData, bodyerr := ioutil.ReadAll(recorder.Body)

	if bodyerr != nil {
		println(bodyerr)
	}

	var returnedUser db.User
	err = json.Unmarshal(bodyData, &returnedUser)

	require.NoError(t, err)
	require.Equal(t, userParams.ID, returnedUser.ID)

}
