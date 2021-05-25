package main

import (
	"fmt"
	"log"
	"time"
)

type UserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email 	 string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponse(user *User) *UserResponse {
	return &UserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}
}

type CreateUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email 	 string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUserResponse struct {
	AccessToken 	string `json:"access_token"`
	UserResponse  	UserResponse `json:"user"`
}

func NewLoginUserResponse(userResponse *UserResponse) *LoginUserResponse {
	tokenMaker, err := NewPasetoMaker(SymmetricKey)
	if err != nil {
		log.Println(fmt.Errorf("cannot create token : %v", err))
	}

	accessToken, err := tokenMaker.CreateToken(userResponse.Username, 15)
	if err != nil {
		log.Println(fmt.Errorf("cannot create token : %v", err))
	}

	return &LoginUserResponse{
		AccessToken: accessToken,
		UserResponse: *userResponse,
	}
}