// Code generated by sqlc. DO NOT EDIT.
// source: messages.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addConversation = `-- name: AddConversation :one
INSERT INTO conversation (
    creator_id,
    recipient_id
) VALUES ($1, $2) RETURNING id, creator_id, recipient_id, created_at, status
`

type AddConversationParams struct {
	CreatorID   uuid.UUID `json:"creator_id"`
	RecipientID uuid.UUID `json:"recipient_id"`
}

func (q *Queries) AddConversation(ctx context.Context, arg AddConversationParams) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, addConversation, arg.CreatorID, arg.RecipientID)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatorID,
		&i.RecipientID,
		&i.CreatedAt,
		&i.Status,
	)
	return i, err
}

const addMessage = `-- name: AddMessage :one
INSERT INTO messages (
    conversation_id,
    sender_id,
    recipient_id,
    message_body
) VALUES ($1, $2, $3, $4) RETURNING id, conversation_id, sender_id, recipient_id, message_body, created_at
`

type AddMessageParams struct {
	ConversationID sql.NullInt32 `json:"conversation_id"`
	SenderID       uuid.UUID     `json:"sender_id"`
	RecipientID    uuid.UUID     `json:"recipient_id"`
	MessageBody    string        `json:"message_body"`
}

func (q *Queries) AddMessage(ctx context.Context, arg AddMessageParams) (Message, error) {
	row := q.db.QueryRowContext(ctx, addMessage,
		arg.ConversationID,
		arg.SenderID,
		arg.RecipientID,
		arg.MessageBody,
	)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.ConversationID,
		&i.SenderID,
		&i.RecipientID,
		&i.MessageBody,
		&i.CreatedAt,
	)
	return i, err
}

const getMessageCommunication = `-- name: GetMessageCommunication :many
SELECT id, conversation_id, sender_id, recipient_id, message_body, created_at FROM messages as m
WHERE m.sender_id = $1
`

func (q *Queries) GetMessageCommunication(ctx context.Context, senderID uuid.UUID) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, getMessageCommunication, senderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ConversationID,
			&i.SenderID,
			&i.RecipientID,
			&i.MessageBody,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}