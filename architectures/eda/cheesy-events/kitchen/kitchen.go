package kitchen

import (
	"fmt"
	"time"

	"cheesy-events/eventbus"
	"cheesy-events/orders"
)

// PizzaPreparedEvent represents the event structure for pizza preparation
type PizzaPreparedEvent struct {
	OrderID string
	Pizza   string
}

// KitchenService represents the service responsible for pizza preparation
type KitchenService struct {
	eventBus       *eventbus.EventBus
	pizzasPrepared int
	maxPizzas      int
}

// NewKitchenService creates a new instance of the kitchen service
func NewKitchenService(eventBus *eventbus.EventBus, maxPizzas int) *KitchenService {
	return &KitchenService{
		eventBus:  eventBus,
		maxPizzas: maxPizzas,
	}
}

// PreparePizza prepares a pizza and publishes a pizza prepared event
func (ks *KitchenService) PreparePizza(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		orderPlacedEvent, ok := event.Data.(orders.OrderPlacedEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		// Check if maximum number of pizzas have been prepared
		if ks.pizzasPrepared >= ks.maxPizzas {
			fmt.Println("Maximum number of pizzas prepared. Cannot prepare more.")
			continue
		}

		// Simulate pizza preparation
		fmt.Printf("Pizza prepared: OrderID=%s, Pizza=%s\n", orderPlacedEvent.OrderID, orderPlacedEvent.Pizza)

		// Create the pizza prepared event
		event := eventbus.Event{
			Type:      "PizzaPrepared",
			Timestamp: time.Now(),
			Data: PizzaPreparedEvent{
				OrderID: orderPlacedEvent.OrderID,
				Pizza:   orderPlacedEvent.Pizza,
			},
		}

		// Publish the event
		ks.eventBus.Publish(event)

		// Increment the number of pizzas prepared
		ks.pizzasPrepared++
	}
}
