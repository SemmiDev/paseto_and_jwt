package main

func(in *LoginUserRequest) Check() (*User, error) {
	getUserByUsername := UserData[in.Username]

	if getUserByUsername == nil {
		return nil, ErrUserNotFound
	}

	getPasswordFromUser := getUserByUsername.Password

	err := CheckPassword(in.Password, getPasswordFromUser)
	if err != nil {
		return nil, ErrUsernameAndPasswordIsWrong
	}

	return getUserByUsername, nil
}