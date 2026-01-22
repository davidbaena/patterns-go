package main

import (
	"cheesy-events/client"
	deliveryDi "cheesy-events/delivery/di"
	kitchenDi "cheesy-events/kitchen/di"
	ordersDi "cheesy-events/orders/di"
	"time"

	"cheesy-events/eventbus"
)

func main() {
	// Create a new event bus
	eventBus := eventbus.NewEventBus()

	kitchenDi.ProvideKitchenHandler(eventBus)
	deliveryDi.ProvideDeliveryHandler(eventBus)
	ordersDi.ProvideOrderHandler(eventBus)

	// Place an order
	newClient := client.NewClient(eventBus)

	tryOrder := client.TryOrder{
		OrderID: "123",
		Pizza:   "Cheese Pizza",
		Customer: client.Customer{
			Name:  "John Doe",
			Phone: "1234567890",
		},
	}

	newClient.SendOrderPlaced(tryOrder)

	time.Sleep(100 * time.Second)
}
