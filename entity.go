package main

import (
	"github.com/google/uuid"
	"time"
)

// UserData for storing the data in memory key-value databases
var UserData = make(map[string] *User)

// User for represent in our user table in database
type User struct {
	ID 			uuid.UUID 	`json:"id"`
	Username 	string 		`json:"username"`
	Password 	string 		`json:"password"`
	FullName 	string 		`json:"full_name"`
	Email 	 	string 		`json:"email"`
	CreatedAt	time.Time 	`json:"created_at"`
}
