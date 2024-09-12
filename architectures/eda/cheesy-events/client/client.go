package client

import (
	"cheesy-events/eventbus"
	"cheesy-events/orders/handler"
	"time"
)

type TryOrder struct {
	OrderID  string
	Pizza    string
	Customer Customer
}

type Customer struct {
	Name  string
	Phone string
}
type Client struct {
	eventBus *eventbus.EventBus
}

func NewClient(eventBus *eventbus.EventBus) Client {
	return Client{eventBus: eventBus}
}

func (p Client) SendOrderPlaced(tryOrder TryOrder) {
	data := handler.OrderPlacedEvent{
		OrderID: tryOrder.OrderID,
		Pizza:   tryOrder.Pizza,
	}

	eventBusEvent := eventbus.Event{
		Type:      "OrderPlaced",
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(eventBusEvent)
}
