package main

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	"time"
)

// JWTMaker is a json web token master
type JWTMaker struct {
	secretKey string
}

// PasetoMaker is a paseto token master
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < 32 {
		return nil, fmt.Errorf("invalid key size must be at least %d character", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

// Maker is an interface for managing token
type Maker interface {
	// CreateToken created a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken checks if token is valid or not
	VerifyToken(token string) (*Payload, error)
}