package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {

	store := NewStore(testDB)

	// create a channel
	// Channels are designed to connect concurrent goroutines
	errChan := make(chan error)
	resultChan := make(chan NewConversationMessageTxResultParams)

	senderID, errSender := uuid.Parse("46b20d28-b6bf-44d7-a0fa-1c9579e23c92")
	recipientID, errRecipient := uuid.Parse("cefadd5e-369f-4be4-b248-f7b456953e06")



	if errSender != nil || errRecipient != nil {
		fmt.Printf("senderError - %v\nrecipientError - %v",errSender.Error(), errRecipient.Error())
		require.NoError(t, errRecipient)
		require.NoError(t, errRecipient)
	}

	println(senderID.String())
	println(recipientID.String())
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