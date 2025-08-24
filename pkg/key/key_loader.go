package key

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

type KeyLoader struct {
	KeyPath string
}

func NewKeyLoader(path string) Key {
	return &KeyLoader{
		KeyPath: path,
	}
}

func (kl *KeyLoader) LoadPrivateKey() ed25519.PrivateKey {
	filePath := fmt.Sprintf("%s/private.pem", kl.KeyPath)
	read, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed load private key: %s", err)
	}

	block, _ := pem.Decode(read)
	parse, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %s", err)
	}

	keyParse := parse.(ed25519.PrivateKey)
	if len(keyParse) != ed25519.PrivateKeySize {
		log.Fatalf("Invalid key size: must be exactly %d characters", ed25519.PrivateKeySize)
	}

	return keyParse
}

func (kl *KeyLoader) LoadPublicKey() ed25519.PublicKey {
	filePath := fmt.Sprintf("%s/public.pem", kl.KeyPath)
	read, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load public key: %s", err)
	}

	block, _ := pem.Decode(read)
	parse, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse public key: %s", err)
	}

	keyParse := parse.(ed25519.PublicKey)
	if len(keyParse) != ed25519.PublicKeySize {
		log.Fatalf("Invalid key size: must be exactly %d characters", ed25519.PublicKeySize)
	}

	return keyParse
}
