-- name: GetMessageCommunication :many
SELECT * FROM messages as m
WHERE m.sender_id = $1;

-- name: AddMessage :one
INSERT INTO messages (
    conversation_id,
    sender_id,
    recipient_id,
    message_body
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: AddConversation :one
INSERT INTO conversation (
    creator_id,
    recipient_id
) VALUES ($1, $2) RETURNING *;

--  name: GetConversation :one
SELECT * FROM conversation As c WHERE c.creator_id = $1;