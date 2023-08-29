package token

import (
	"log"
	"os"
	"time"

	errs "github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-Mayhul-Jindal/errors"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// pasetoMaker is a PASETO token maker
type pasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new pasetoMaker
func NewPasetoMaker(symmetricKey string) Maker {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		log.Fatalf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
		os.Exit(1)
	}

	return &pasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
}

// CreateToken creates a new token for a specific username and duration
func (maker *pasetoMaker) CreateToken(userID int64, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *pasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, errs.ErrorInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
