package domain

import (
	"errors"
)

type User struct {
	Email    Email
	Password Password
}

func NewUser(email Email, password Password) (*User, error) {
	if email == (Email{}) || password == (Password{}) {
		return nil, errors.New("invalid email or password")
	}
	return &User{Email: email, Password: password}, nil
}
