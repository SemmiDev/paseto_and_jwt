package main

import (
	"github.com/google/uuid"
	"time"
)

// Payload container the payload data of the token
type Payload struct {
	ID 			uuid.UUID 	`json:"id"`
	Username 	string 		`json:"username"`
	IssuedAt 	time.Time  	`json:"issued_at"`
	ExpiredAt 	time.Time 	`json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

// Valid checks if the token valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}