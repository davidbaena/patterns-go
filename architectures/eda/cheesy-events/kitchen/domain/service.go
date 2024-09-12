package domain

import (
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"
)

type PizzaRepository interface {
	CreatePizza(pizza Pizza) (int64, error)
	GetPizza(id int) (Pizza, error)
	UpdatePizza(pizza Pizza) error
	DeletePizza(id int) error
}

type Producer interface {
	SendPizzaBeingPrepared(pizza Pizza)
	SendPizzaPrepared(pizza Pizza)
}

var logger = logrus.NewLogger("KitchenService")

type KitchenService struct {
	producer       Producer
	repository     PizzaRepository
	pizzasPrepared int
	maxPizzas      int
}

func NewKitchenService(producer Producer, repository PizzaRepository, maxPizzas int) KitchenService {
	return KitchenService{
		producer:   producer,
		repository: repository,
		maxPizzas:  maxPizzas,
	}
}

func (ks KitchenService) PreparePizza(pizza Pizza) {
	if ks.pizzasPrepared >= ks.maxPizzas {
		logger.Error("Maximum number of pizzas prepared. Cannot prepare more.")
		return
	}

	// 1 - Persist the pizza
	if _, err := ks.repository.CreatePizza(pizza); err != nil {
		logger.WithFields(log.Fields{
			"order_id": pizza.OrderId,
			"pizza":    pizza.Name,
			"error":    err,
		}).Error("Failed to store pizza")
		return
	}

	pizza.nextState() // to CollectingIngredients
	if err := ks.repository.UpdatePizza(pizza); err != nil {
		logger.WithFields(log.Fields{
			"order_id": pizza.OrderId,
			"pizza":    pizza.Name,
			"error":    err,
		}).Error("Failed to update pizza")
		return
	}

	logger.WithFields(log.Fields{
		"order_id": pizza.OrderId,
		"pizza":    pizza.Name,
	}).Info("Collecting Ingredients...")

	pizza.nextState() // to OvenBaking

	if err := ks.repository.UpdatePizza(pizza); err != nil {
		logger.WithFields(log.Fields{
			"order_id": pizza.OrderId,
			"pizza":    pizza.Name,
			"error":    err,
		}).Error("Failed to update pizza")
		return
	}

	// 2 - Send the pizza being prepared event
	ks.producer.SendPizzaBeingPrepared(pizza)

	logger.WithFields(log.Fields{
		"order_id": pizza.OrderId,
		"pizza":    pizza.Name,
	}).Info("Preparing pizza...")

	time.Sleep(1 * time.Second)

	// 3 - Pizza prepared
	pizza.nextState() //PizzaPrepared
	if err := ks.repository.UpdatePizza(pizza); err != nil {
		logger.WithFields(log.Fields{
			"order_id": pizza.OrderId,
			"pizza":    pizza.Name,
			"error":    err,
		}).Error("Failed to update pizza")
		return
	}

	logger.WithFields(log.Fields{
		"order_id": pizza.OrderId,
		"pizza":    pizza.Name,
		"status":   pizza.Status,
	}).Info("PizzaName prepared")

	ks.producer.SendPizzaPrepared(pizza)

	ks.pizzasPrepared++
}
