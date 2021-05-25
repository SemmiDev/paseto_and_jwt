package main

import (
	"github.com/google/uuid"
	"time"
)

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	FullName 	string `json:"full_name"`
	Email 	 	string `json:"email"`
}

func (user *CreateUserRequest) createNewUser() *CreateUserResponse {
	id, _ := uuid.NewRandom()
	hashPass, _ := HashPassword(user.Password)

	newUser := User{
		ID:        id,
		Username:  user.Username,
		Password:  hashPass,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: time.Now(),
	}

	// storing user in memory db
	UserData[user.Username] = &newUser

	return &CreateUserResponse{
		Username:  newUser.Username,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}
}