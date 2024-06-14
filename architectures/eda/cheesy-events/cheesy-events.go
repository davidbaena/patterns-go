package main

import (
	"time"

	"cheesy-events/delivery"
	"cheesy-events/eventbus"
	"cheesy-events/kitchen"
	"cheesy-events/orders"
)

func main() {
	// Create a new event bus
	eventBus := eventbus.NewEventBus()

	// Create the services
	orderService := orders.NewOrderService(eventBus)
	kitchenService := kitchen.NewKitchenService(eventBus, 2)
	deliveryService := delivery.NewDeliveryService(eventBus)

	// Create channels for the events
	orderPlacedChan := make(chan eventbus.Event)
	pizzaPreparedChan := make(chan eventbus.Event)

	// Subscribe the channels to the events
	eventBus.Subscribe("OrderPlaced", orderPlacedChan)
	eventBus.Subscribe("PizzaPrepared", pizzaPreparedChan)

	// Start the event handlers
	go kitchenService.PreparePizza(orderPlacedChan)
	go deliveryService.DeliverPizza(pizzaPreparedChan)

	// Place an order
	orderService.PlaceOrder("1", "Margherita")
	orderService.PlaceOrder("1", "Margherita")
	orderService.PlaceOrder("1", "Margherita")

	time.Sleep(2 * time.Second)
}
