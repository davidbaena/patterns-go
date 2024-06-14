package orders

import (
	"fmt"
	"time"

	"cheesy-events/eventbus"
)

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

// OrderPlacedEvent represents the event structure for order placement
type OrderPlacedEvent struct {
	OrderID string
	Pizza   string
}

// PlaceOrder places a new order and publishes an order placed event
func (os *OrderService) PlaceOrder(orderID, pizza string) {
	// Simulate order placement
	fmt.Printf("Order placed: OrderID=%s, Pizza=%s\n", orderID, pizza)

	// Create the order placed event
	event := eventbus.Event{
		Type:      "OrderPlaced",
		Timestamp: time.Now(),
		Data: OrderPlacedEvent{
			OrderID: orderID,
			Pizza:   pizza,
		},
	}

	// Publish the event
	os.eventBus.Publish(event)
}
