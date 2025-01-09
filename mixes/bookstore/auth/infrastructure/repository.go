package infrastructure

import (
	"sync"

	"bookstore/auth/domain"
)

type InMemoryUserRepository struct {
	users map[string]*domain.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.Email.String()] = user
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(email domain.Email) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[email.String()]
	if !exists {
		return nil, nil
	}
	return user, nil
}
