package adapter

import (
	"cheesy-events/eventbus"
	"cheesy-events/orders/domain"
	"time"
)

// OrderAcceptedEvent represents the event structure for order acceptance
type OrderAcceptedEvent struct {
	OrderID string
	Pizza   string
}

type OrderProducer struct {
	eventBus *eventbus.EventBus
}

func NewOrderProducer(eventBus *eventbus.EventBus) OrderProducer {
	return OrderProducer{eventBus: eventBus}
}

func (p OrderProducer) SendOrderAccepted(order domain.Order) {
	data := OrderAcceptedEvent{
		OrderID: order.ID,
		Pizza:   order.Pizza,
	}

	eventBusEvent := eventbus.Event{
		Type:      "OrderAccepted",
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(eventBusEvent)
}
