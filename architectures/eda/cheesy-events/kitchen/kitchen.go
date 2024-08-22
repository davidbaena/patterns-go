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

// PizzaBeingPreparedEvent represents the event structure for pizza preparation status
type PizzaBeingPreparedEvent struct {
	OrderID string
	Status  string
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
		fmt.Printf("Pizza being prepared: OrderID=%s, Pizza=%s\n", orderPlacedEvent.OrderID, orderPlacedEvent.Pizza)

		// Create the pizza being prepared event
		preparingEvent := eventbus.Event{
			Type:      "PizzaBeingPrepared",
			Timestamp: time.Now(),
			Data: PizzaBeingPreparedEvent{
				OrderID: orderPlacedEvent.OrderID,
				Status:  "Pizza Being Prepared",
			},
		}
		ks.PrintEvent(preparingEvent)
		ks.eventBus.Publish(preparingEvent)

		fmt.Println("Preparing pizza....")
		time.Sleep(1 * time.Second)
		// Create the pizza prepared event
		event := eventbus.Event{
			Type:      "PizzaPrepared",
			Timestamp: time.Now(),
			Data: PizzaPreparedEvent{
				OrderID: orderPlacedEvent.OrderID,
				Pizza:   orderPlacedEvent.Pizza,
			},
		}
		ks.PrintEvent(event)
		ks.eventBus.Publish(event)

		// Increment the number of pizzas prepared
		ks.pizzasPrepared++
	}
}

func (ks *KitchenService) PrintEvent(event eventbus.Event) {
	fmt.Printf("Publish Event: %s\n", event.Type)
}
