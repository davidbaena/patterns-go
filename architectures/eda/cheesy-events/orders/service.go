package orders

import (
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"

	"cheesy-events/eventbus"
)

var logger = logrus.NewLogger("orders")

// OrderService represents the service responsible for order placement
type OrderService struct {
	eventBus *eventbus.EventBus
}

// NewOrderService creates a new instance of the order service
func NewOrderService(eventBus *eventbus.EventBus) *OrderService {
	return &OrderService{
		eventBus: eventBus,
	}
}

// OrderAcceptedEvent represents the event structure for order acceptance
type OrderAcceptedEvent struct {
	OrderID string
	Pizza   string
	Status  string
}

// OrderPlacedEvent represents the event structure for order placement
type OrderPlacedEvent struct {
	OrderID string
	Pizza   string
	Status  string
}

// PlaceOrder places a new order and publishes an order placed event
func (os *OrderService) PlaceOrder(orderID, pizza string) {
	// Simulate order placement
	logger.WithFields(log.Fields{
		"order_id": orderID,
		"pizza":    pizza,
	}).Info("Order placed")

	// Create the order placed event
	event := eventbus.Event{
		Type:      "OrderPlaced",
		Timestamp: time.Now(),
		Data: OrderPlacedEvent{
			OrderID: orderID,
			Pizza:   pizza,
			Status:  "Order Placed",
		},
	}

	// Publish the event
	os.eventBus.Publish(event)

	logger.WithFields(log.Fields{
		"order_id": orderID,
		"pizza":    pizza,
	}).Info("Checking Fraud system....")

	time.Sleep(1 * time.Second)

	// Simulate order acceptance
	acceptedEvent := eventbus.Event{
		Type:      "OrderAccepted",
		Timestamp: time.Now(),
		Data: OrderAcceptedEvent{
			OrderID: orderID,
			Pizza:   pizza,
			Status:  "Order Accepted",
		},
	}
	os.eventBus.Publish(acceptedEvent)
}
