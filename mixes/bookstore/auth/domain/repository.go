package domain

type UserRepository interface {
	Save(user *User) error
	FindByEmail(email Email) (*User, error)
}
