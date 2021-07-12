package authentication

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorExpiredToken = errors.New("token has expired")
	ErrorInvalidToken = errors.New("token invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"created_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	return &Payload{ID: tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration)}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrorExpiredToken
	}

	return nil
}
