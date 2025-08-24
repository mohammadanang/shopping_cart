package middleware

import (
	"github.com/mohammadanang/shopping-cart/pkg/config"
	"github.com/mohammadanang/shopping-cart/pkg/token"
	"github.com/mohammadanang/shopping-cart/pkg/token/asymmetric"
)

const (
	authorizationHeader     = "authorization"
	authorizationBearerType = "bearer"
	authorizationPayload    = "authorization_payload"
)

type Middleware struct {
	tokenMaker token.Maker
}

func NewMiddleware(conf *config.Config) *Middleware {
	return &Middleware{
		tokenMaker: asymmetric.NewMaker(conf.PublicKey, conf.PrivateKey),
	}
}
