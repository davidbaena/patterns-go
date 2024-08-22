package tracking

import (
	"cheesy-events/delivery"
	"cheesy-events/eventbus"
	"cheesy-events/kitchen"
	"cheesy-events/orders"
	"fmt"
)

// OrderTrackingService represents the service responsible for tracking order status
type OrderTrackingService struct {
	eventBus *eventbus.EventBus
}

// NewOrderTrackingService creates a new instance of the order tracking service
func NewOrderTrackingService(eventBus *eventbus.EventBus) *OrderTrackingService {
	return &OrderTrackingService{
		eventBus: eventBus,
	}
}

// TrackOrderStatus listens for order status events and updates the order status
func (ots *OrderTrackingService) TrackOrderStatus(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		switch e := event.Data.(type) {
		case orders.OrderAcceptedEvent:
			fmt.Printf("Order accepted: OrderID=%s, Status=%s\n", e.OrderID, e.Status)
		case kitchen.PizzaBeingPreparedEvent:
			fmt.Printf("Pizza being prepared: OrderID=%s, Status=%s\n", e.OrderID, e.Status)
		case delivery.OutForDeliveryEvent:
			fmt.Printf("Out for delivery: OrderID=%s, Status=%s\n", e.OrderID, e.Status)
		case delivery.DeliveredEvent:
			fmt.Printf("Delivered: OrderID=%s, Status=%s\n", e.OrderID, e.Status)
		default:
			fmt.Println("Unknown event type")
		}
	}
}
