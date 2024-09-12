package domain

import (
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"

	"cheesy-events/orders"
)

// PizzaRepository defines the methods for managing pizzas
type PizzaRepository interface {
	CreatePizza(pizza Pizza) (int64, error)
	GetPizza(id int) (Pizza, error)
	UpdatePizza(pizza Pizza) error
	DeletePizza(id int) error
}

// Create an interface for the producer
type Producer interface {
	SendPizzaBeingPrepared(pizza Pizza)
	SendPizzaPrepared(pizza Pizza)
}

var logger = logrus.NewLogger("KitchenService")

type KitchenService struct {
	producer       Producer
	pizzasPrepared int
	maxPizzas      int
	repository     PizzaRepository
}

func NewKitchenService(producer Producer, repository PizzaRepository, maxPizzas int) KitchenService {
	return KitchenService{
		producer:   producer,
		maxPizzas:  maxPizzas,
		repository: repository,
	}
}

func (ks *KitchenService) PreparePizza(orderAcceptedEvent orders.OrderAcceptedEvent) {
	if ks.pizzasPrepared >= ks.maxPizzas {
		logger.Error("Maximum number of pizzas prepared. Cannot prepare more.")
		return
	}

	pizza := Pizza{
		OrderId: orderAcceptedEvent.OrderID,
		Name:    orderAcceptedEvent.Pizza,
		Status:  CollectingIngredients,
	}

	logger.WithFields(log.Fields{
		"order_id": pizza.OrderId,
		"pizza":    pizza.Name,
	}).Info("Collecting Ingredients...")

	pizza.nextState()

	ks.producer.SendPizzaBeingPrepared(pizza)

	logger.WithFields(log.Fields{
		"order_id": pizza.OrderId,
		"pizza":    pizza.Name,
	}).Info("Preparing pizza...")

	time.Sleep(1 * time.Second)

	pizza.nextState()

	logger.WithFields(log.Fields{
		"order_id": pizza.OrderId,
		"pizza":    pizza.Name,
		"status":   pizza.Status,
	}).Info("PizzaName prepared")

	ks.producer.SendPizzaPrepared(pizza)

	ks.pizzasPrepared++
}
