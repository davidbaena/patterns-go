package adapter

import (
	"cheesy-events/kitchen/domain"
	"errors"
	"sync"
)

type RepositoryDummy struct {
	mu    sync.Mutex
	store map[string]domain.Pizza
}

func NewRepositoryDummy() RepositoryDummy {
	return RepositoryDummy{
		store: make(map[string]domain.Pizza),
	}
}

func (r *RepositoryDummy) CreatePizza(pizza domain.Pizza) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[pizza.ID] = pizza
	return 1, nil
}

func (r *RepositoryDummy) GetPizza(id int) (domain.Pizza, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	pizza, exists := r.store[string(id)]
	if !exists {
		return domain.Pizza{}, errors.New("pizza not found")
	}
	return pizza, nil
}

func (r *RepositoryDummy) UpdatePizza(pizza domain.Pizza) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.store[pizza.ID]
	if !exists {
		return errors.New("pizza not found")
	}
	r.store[pizza.ID] = pizza
	return nil
}

func (r *RepositoryDummy) DeletePizza(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.store, string(id))
	return nil
}
