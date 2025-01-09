package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hashedValue string
}

func NewPassword(value string) (Password, error) {
	if len(value) < 8 {
		return Password{}, errors.New("password must be at least 8 characters long")
	}
	hashedValue, err := hashPassword(value)
	if err != nil {
		return Password{}, err
	}
	return Password{hashedValue: hashedValue}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p Password) String() string {
	return p.hashedValue
}
