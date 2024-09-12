package adapter

import (
	"cheesy-events/delivery/domain"
	"errors"
	"sync"
)

type DeliveryRepositoryDummy struct {
	mu    sync.Mutex
	store map[string]domain.Delivery
}

func NewDeliveryRepositoryDummy() DeliveryRepositoryDummy {
	return DeliveryRepositoryDummy{
		store: make(map[string]domain.Delivery),
	}
}

func (r *DeliveryRepositoryDummy) CreateDelivery(delivery domain.Delivery) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[delivery.ID] = delivery
	return 1, nil
}

func (r *DeliveryRepositoryDummy) GetDelivery(id string) (domain.Delivery, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delivery, exists := r.store[id]
	if !exists {
		return domain.Delivery{}, errors.New("delivery not found")
	}
	return delivery, nil
}

func (r *DeliveryRepositoryDummy) UpdateDelivery(delivery domain.Delivery) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.store[delivery.ID]
	if !exists {
		return errors.New("delivery not found")
	}
	r.store[delivery.ID] = delivery
	return nil
}

func (r *DeliveryRepositoryDummy) DeleteDelivery(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.store, id)
	return nil
}
