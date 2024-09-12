package main

import (
	deliveryDi "cheesy-events/delivery/di"
	kitchenDi "cheesy-events/kitchen/di"
	"time"

	"cheesy-events/eventbus"
	"cheesy-events/orders"
)

func main() {
	// Create a new event bus
	eventBus := eventbus.NewEventBus()

	// Create the services
	orderService := orders.NewOrderService(eventBus)

	kitchenDi.ProvideKitchenHandler(eventBus)
	deliveryDi.ProvideDeliveryHandler(eventBus)

	// Place an order
	orderService.PlaceOrder("1", "Margherita")
	//orderService.PlaceOrder("2", "BBQ Chicken")
	//orderService.PlaceOrder("3", "Pepperoni")
	//orderService.PlaceOrder("4", "Vegetarian")

	time.Sleep(100 * time.Second)
}
