package main

import (
	"cheesy-events/tracking"
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
	trackingService := tracking.NewOrderTrackingService(eventBus)

	// Create channels for the events
	orderPlacedChan := make(chan eventbus.Event)
	orderAcceptedChan := make(chan eventbus.Event)
	pizzaBeingPreparedChan := make(chan eventbus.Event)
	pizzaPreparedChan := make(chan eventbus.Event)
	outForDeliveryChan := make(chan eventbus.Event)
	deliveredChan := make(chan eventbus.Event)

	// Subscribe the channels to the events
	eventBus.Subscribe("OrderPlaced", orderPlacedChan)
	eventBus.Subscribe("OrderAccepted", orderAcceptedChan)
	eventBus.Subscribe("PizzaBeingPrepared", pizzaBeingPreparedChan)
	eventBus.Subscribe("PizzaPrepared", pizzaPreparedChan)
	eventBus.Subscribe("OutForDelivery", outForDeliveryChan)
	eventBus.Subscribe("Delivered", deliveredChan)

	// Start the event handlers
	go kitchenService.PreparePizza(orderPlacedChan)
	go deliveryService.DeliverPizza(pizzaPreparedChan)
	go trackingService.TrackOrderStatus(orderAcceptedChan)
	go trackingService.TrackOrderStatus(pizzaBeingPreparedChan)
	go trackingService.TrackOrderStatus(outForDeliveryChan)
	go trackingService.TrackOrderStatus(deliveredChan)

	// Place an order
	orderService.PlaceOrder("1", "Margherita")
	orderService.PlaceOrder("2", "BBQ Chicken")
	orderService.PlaceOrder("3", "Pepperoni")
	orderService.PlaceOrder("4", "Vegetarian")

	time.Sleep(2 * time.Second)
}
