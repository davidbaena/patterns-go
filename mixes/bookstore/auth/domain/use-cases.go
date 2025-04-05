package domain

import (
	"errors"
)

// RegisterUser registers a new user.
func (s *AuthService) RegisterUser(email Email, password string) error {
	// Check if user already exists
	_, err := s.userRepository.FindByEmail(email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash the password
	hashedPassword, err := s.passwordHasher.Hash(password)
	if err != nil {
		return err
	}

	// Create the user
	user := User{
		Email: email,
		Password: Password{
			hashedValue: hashedPassword,
		},
	}
	return s.userRepository.Save(&user)
}

// AuthenticateUser authenticates a user.
func (s *AuthService) AuthenticateUser(email Email, password string) (User, error) {
	// Get the user by email
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return User{}, errors.New("user not found")
	}

	// Compare the password
	err = s.passwordHasher.Compare(user.Password.hashedValue, password)
	if err != nil {
		return User{}, errors.New("invalid password")
	}

	return *user, nil
}
