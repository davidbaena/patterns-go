package delivery

import (
	"fmt"

	"cheesy-events/eventbus"
	"cheesy-events/kitchen"
)

// DeliveryService represents the service responsible for pizza delivery
type DeliveryService struct {
	eventBus *eventbus.EventBus
}

// NewDeliveryService creates a new instance of the delivery service
func NewDeliveryService(eventBus *eventbus.EventBus) *DeliveryService {
	return &DeliveryService{
		eventBus: eventBus,
	}
}

// DeliverPizza delivers a pizza when a pizza prepared event is received
func (ds *DeliveryService) DeliverPizza(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		pizzaPreparedEvent, ok := event.Data.(kitchen.PizzaPreparedEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		// Simulate pizza delivery
		fmt.Printf("Pizza delivered: OrderID=%s, Pizza=%s\n", pizzaPreparedEvent.OrderID, pizzaPreparedEvent.Pizza)
	}
}
