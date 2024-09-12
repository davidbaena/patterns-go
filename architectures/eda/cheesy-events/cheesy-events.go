package main

import (
	"cheesy-events/kitchen/adapter"
	"cheesy-events/kitchen/domain"
	"cheesy-events/kitchen/handler"
	"cheesy-events/sqlclient"
	"cheesy-events/tracking"
	"time"

	"cheesy-events/delivery"
	"cheesy-events/eventbus"
	"cheesy-events/orders"
)

func main() {
	// Create a new event bus
	eventBus := eventbus.NewEventBus()

	// Create the services
	orderService := orders.NewOrderService(eventBus)
	sql := sqlclient.NewSqClient()
	producer := adapter.NewProducer(eventBus)
	kitchenService := domain.NewKitchenService(producer, sql, 2)
	handler.NewEventHandler(eventBus, kitchenService)
	delivery.NewDeliveryService(eventBus)
	tracking.NewOrderTrackingService(eventBus)

	// Place an order
	orderService.PlaceOrder("1", "Margherita")
	orderService.PlaceOrder("2", "BBQ Chicken")
	orderService.PlaceOrder("3", "Pepperoni")
	//orderService.PlaceOrder("4", "Vegetarian")

	time.Sleep(2 * time.Second)
}
