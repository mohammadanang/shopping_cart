package key

import "crypto/ed25519"

type Key interface {
	LoadPrivateKey() ed25519.PrivateKey
	LoadPublicKey() ed25519.PublicKey
}
