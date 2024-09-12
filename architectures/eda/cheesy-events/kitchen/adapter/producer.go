package adapter

import (
	"cheesy-events/eventbus"
	"cheesy-events/kitchen/domain"
	"time"
)

// PizzaPreparedEvent represents the event structure for pizza preparation
type PizzaPreparedEvent struct {
	eventbus.Event
	OrderID string
	Pizza   string
}

// PizzaBeingPreparedEvent represents the event structure for pizza preparation status
type PizzaBeingPreparedEvent struct {
	OrderID string
	Status  string
}

type Producer struct {
	eventBus *eventbus.EventBus
}

func NewProducer(eventBus *eventbus.EventBus) Producer {
	return Producer{eventBus: eventBus}
}

func (p Producer) SendPizzaBeingPrepared(pizza domain.Pizza) {
	data := PizzaBeingPreparedEvent{
		OrderID: pizza.OrderId,
		Status:  string(pizza.Status),
	}

	event := eventbus.Event{
		Type:      "PizzaBeingPreparedEvent",
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(event)
}

func (p Producer) SendPizzaPrepared(pizza domain.Pizza) {

	data := PizzaPreparedEvent{
		OrderID: pizza.OrderId,
		Pizza:   pizza.Name,
	}

	event := eventbus.Event{
		Type:      "PizzaPreparedEvent",
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(event)
}
