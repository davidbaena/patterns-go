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

// OutForDeliveryEvent represents the event structure for out for delivery status
type OutForDeliveryEvent struct {
	OrderID string
	Status  string
}

// DeliveredEvent represents the event structure for delivered status
type DeliveredEvent struct {
	OrderID string
	Status  string
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

		// Create the out for delivery event
		outForDeliveryEvent := eventbus.Event{
			Type:      "OutForDelivery",
			Timestamp: time.Now(),
			Data: OutForDeliveryEvent{
				OrderID: pizzaPreparedEvent.OrderID,
				Status:  "Out For Delivery",
			},
		}
		ds.PrintEvent(outForDeliveryEvent)
		ds.eventBus.Publish(outForDeliveryEvent)

		fmt.Println("Delivering pizza....")
		time.Sleep(1 * time.Second)

		// Create the delivered event
		deliveredEvent := eventbus.Event{
			Type:      "Delivered",
			Timestamp: time.Now(),
			Data: DeliveredEvent{
				OrderID: pizzaPreparedEvent.OrderID,
				Status:  "Delivered",
			},
		}
		ds.PrintEvent(deliveredEvent)
		ds.eventBus.Publish(deliveredEvent)
	}
}

func (ds *DeliveryService) PrintEvent(event eventbus.Event) {
	fmt.Printf("Publish Event: %s\n", event.Type)
}
