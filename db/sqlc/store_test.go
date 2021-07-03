package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {

	store := NewStore(testDB)
	errChan := make(chan error)
	resultChan := make(chan NewConversationMessageTxResultParams)

	senderID, recipientID := uuid.New(), uuid.New()

	newUserUUIDs := [2]uuid.UUID{senderID, recipientID}

	inserChan := make(chan error, 2)

	go func() {
		for i := range [2]int{} {
			_, insertErr := store.CreateUser(context.Background(), CreateUserParams{
				ID: newUserUUIDs[i],
				FullName: fmt.Sprintf("test user %v", i + 1),
				Contact: "01111111111",
				Dog: sql.NullInt32{Int32: 1, Valid: true},
				Address: fmt.Sprintf("test address %v", i + 1),
				City: fmt.Sprintf("city %v", i + 1),
				PostCode: "D1 1AA",
				Longitude: "1268327832",
				Latitude: "473987493",
			})

			inserChan <- insertErr
		}

		}()

		err1,err2 := <-inserChan,<-inserChan


		require.NoError(t, err1)
		require.NoError(t, err2)


	go func() {
		result, err := store.messageTx(context.Background(),NewConversationMessageTxParams{
			SenderID: senderID,
			RecipientID: recipientID,
			MessageBody: "hello world",
		})

		// the goroutine is considered the sender here as a chanel is being sent some data <-
		errChan <- err
		resultChan <- result
	}()

	// the err variable declared here is the reciever, since it is recieving data from a chanel. Recievings are blockers
	// therefore the code will only proceed once it has received

	err := <- errChan
	require.NoError(t, err)

	result := <- resultChan
	require.NotEmpty(t, result)

	conversation := result.Conversation
	require.NotEmpty(t, conversation)

	message := result.Message
	require.NotEmpty(t, message)

	require.Equal(t, message.RecipientID, recipientID )
	require.Equal(t, message.MessageBody, "hello world")

}