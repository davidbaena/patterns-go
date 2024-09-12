package domain

import (
	"cheesy-events/eventbus"
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"
)

var logger = logrus.NewLogger("orders")

type OrderService struct {
	eventBus *eventbus.EventBus
	producer OrderProducer
}
type OrderProducer interface {
	SendOrderAccepted(order Order)
}

func NewOrderService(eventBus *eventbus.EventBus, producer OrderProducer) OrderService {
	return OrderService{
		eventBus: eventBus,
		producer: producer,
	}
}

func (os *OrderService) PlaceOrder(order Order) {
	logger.WithFields(log.Fields{
		"order_id": order.ID,
		"pizza":    order.Pizza,
	}).Info("Order placed")

	logger.WithFields(log.Fields{
		"order_id": order.ID,
		"pizza":    order.Pizza,
	}).Info("Checking Fraud system....")

	time.Sleep(1 * time.Second)

	os.producer.SendOrderAccepted(order)
}
