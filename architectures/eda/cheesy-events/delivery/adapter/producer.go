package adapter

import (
	"cheesy-events/delivery/domain"
	"cheesy-events/eventbus"
	"time"
)

type DeliveryProducer struct {
	eventBus *eventbus.EventBus
}

func NewDeliveryProducer(eventBus *eventbus.EventBus) DeliveryProducer {
	return DeliveryProducer{eventBus: eventBus}
}

func (p DeliveryProducer) SendDeliveryStarted(delivery domain.Delivery) {
	data := map[string]string{
		"OrderID": delivery.OrderID,
		"Status":  string(delivery.Status),
	}

	event := eventbus.Event{
		Type:      "DeliveryStarted",
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(event)
}

func (p DeliveryProducer) SendDeliveryCompleted(delivery domain.Delivery) {
	data := map[string]string{
		"OrderID": delivery.OrderID,
		"Status":  string(delivery.Status),
	}

	event := eventbus.Event{
		Type:      "DeliveryCompleted",
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(event)
}
