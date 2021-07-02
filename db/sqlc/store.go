package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)


type Store struct {
	*Queries
	db *sql.DB
}


func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// lowerCase first letter functions are considered private the package they are in
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	// instead of passing in a DB context, we now pass in the database transaction context
	query := New(tx)
	err = fn(query)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil { return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr ) }

		return err
	}

	return tx.Commit()

}

type NewConversationMessageTxParams struct{
	SenderID uuid.UUID `json:"sender_id"`
	RecipientID uuid.UUID `json:"recipient_id"`
	MessageBody string `json:"message_body"`
}

type NewConversationMessageTxResultParams struct{
	Message Message `json:"message"`
	Conversation Conversation `json:"conversation"`
}


func (store * Store) messageTx(ctx context.Context, arg NewConversationMessageTxParams) (NewConversationMessageTxResultParams, error) {
	var result NewConversationMessageTxResultParams

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Conversation, err = q.AddConversation(ctx, AddConversationParams{CreatorID: arg.SenderID, RecipientID: arg.RecipientID})

		if err != nil {
			return err
		}

		result.Message, err = q.AddMessage(ctx,AddMessageParams{
			ConversationID: sql.NullInt32{Int32: result.Conversation.ID, Valid: true},
			SenderID: arg.SenderID,
			RecipientID: arg.RecipientID,
			MessageBody: arg.MessageBody})

		if err != nil {
			return err
		}

		return nil

	})

	return result, err
}





