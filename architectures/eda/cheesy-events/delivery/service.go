package delivery

import (
	"fmt"
	"time"

	"cheesy-events/eventbus"
	"cheesy-events/kitchen"
)

// DeliveryInProgressEvent represents the event structure for pizza delivery in progress
type PizzaDeliveredEvent struct {
	OrderID string
	Pizza   string
}

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

		deliveryInProgressEvent := eventbus.Event{
			Type:      "PizzaPrepared",
			Timestamp: time.Now(),
			Data: PizzaDeliveredEvent{
				OrderID: pizzaPreparedEvent.OrderID,
				Pizza:   pizzaPreparedEvent.Pizza,
			},
		}

		ds.PrintEvent(deliveryInProgressEvent)
	}
}

func (ds *DeliveryService) PrintEvent(event eventbus.Event) {
	fmt.Printf("Publish Event: %s\n", event.Type)
}
