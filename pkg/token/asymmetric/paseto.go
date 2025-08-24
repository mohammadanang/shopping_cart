package asymmetric

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mohammadanang/shopping-cart/pkg/token"
	"github.com/o1egl/paseto"
)

type pasetoMaker struct {
	paseto     *paseto.V2
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

func NewMaker(
	publicKey ed25519.PublicKey,
	privateKey ed25519.PrivateKey,
) token.Maker {
	pasetoV2 := paseto.NewV2()
	maker := &pasetoMaker{
		paseto:     pasetoV2,
		publicKey:  publicKey,
		privateKey: privateKey,
	}

	return maker
}

func (maker *pasetoMaker) CreateToken(username string, duration time.Duration) (string, *token.Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", nil, err
	}

	payload := token.NewPayload(username, duration, tokenId)
	token, err := maker.paseto.Sign(maker.privateKey, payload, nil)
	return token, payload, err
}

func (maker *pasetoMaker) VerifyToken(tokenValue string) (*token.Payload, error) {
	var payload token.Payload
	err := maker.paseto.Verify(tokenValue, maker.publicKey, &payload, nil)
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
