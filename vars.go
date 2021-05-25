package main

import "errors"

// Err responses for errors
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
	ErrUserNotFound = errors.New("user not found")
	ErrUsernameAndPasswordIsWrong = errors.New("username/password is wrong")
)

// Constant value for token usage
const (
	SymmetricKey      = "12345678901234567890123456789012"
	minSecretKeySize  = 32
)