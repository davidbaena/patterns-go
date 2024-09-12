package domain

import (
	"cheesy-events/kitchen/adapter"
	"cheesy-events/sqlclient"
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"

	"cheesy-events/orders"
)

var logger = logrus.NewLogger("KitchenService")

// PizzaPreparedEvent represents the event structure for pizza preparation
type PizzaPreparedEvent struct {
	OrderID string
	Pizza   string
}

// PizzaBeingPreparedEvent represents the event structure for pizza preparation status
type PizzaBeingPreparedEvent struct {
	OrderID string
	Status  string
}

type KitchenService struct {
	producer       adapter.Producer
	pizzasPrepared int
	maxPizzas      int
	sqClient       sqlclient.SQLClient
}

func NewKitchenService(producer adapter.Producer, sqClient sqlclient.SQLClient, maxPizzas int) KitchenService {
	return KitchenService{
		producer:  producer,
		maxPizzas: maxPizzas,
		sqClient:  sqClient,
	}
}

func (ks *KitchenService) PreparePizza(orderAcceptedEvent orders.OrderAcceptedEvent) {
	if ks.pizzasPrepared >= ks.maxPizzas {
		logger.Error("Maximum number of pizzas prepared. Cannot prepare more.")
		return
	}

	preparingEvent := PizzaBeingPreparedEvent{
		OrderID: orderAcceptedEvent.OrderID,
		Status:  "Pizza Being Prepared",
	}
	ks.producer.ProduceEvent("PizzaBeingPrepared", preparingEvent)

	logger.WithFields(log.Fields{
		"order_id": orderAcceptedEvent.OrderID,
		"pizza":    orderAcceptedEvent.Pizza,
	}).Info("Preparing pizza...")

	time.Sleep(1 * time.Second)

	preparedEvent := PizzaPreparedEvent{
		OrderID: orderAcceptedEvent.OrderID,
		Pizza:   orderAcceptedEvent.Pizza,
	}
	logger.WithFields(log.Fields{
		"order_id": orderAcceptedEvent.OrderID,
		"pizza":    orderAcceptedEvent.Pizza,
	}).Info("Pizza prepared")

	ks.producer.ProduceEvent("PizzaPrepared", preparedEvent)

	ks.pizzasPrepared++
}
