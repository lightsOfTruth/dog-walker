package authentication

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const MIN_SECRET_KEY_SIZE = 32

type JWTCreator struct {
	secretKey string
}

func NewJWTCreator(secretKey string) (TokenCreator, error) {
	if len(secretKey) < MIN_SECRET_KEY_SIZE {
		return nil, fmt.Errorf("invalid key size")
	}

	return &JWTCreator{secretKey}, nil
}

func (jwtCreator *JWTCreator) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	unix := payload.IssuedAt.Unix()
	unix2 := payload.ExpiredAt.Unix()

	fmt.Printf("%v %v", unix, unix2)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(jwtCreator.secretKey))

}

func (jwtCreator *JWTCreator) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}

		return []byte(jwtCreator.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrorInvalidToken
	}

	return payload, nil
}
