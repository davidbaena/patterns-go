package adapter

import (
	"cheesy-events/eventbus"
	"time"
)

type Producer struct {
	eventBus *eventbus.EventBus
}

func NewProducer(eventBus *eventbus.EventBus) Producer {
	return Producer{eventBus: eventBus}
}

func (p *Producer) ProduceEvent(eventType string, data interface{}) {
	event := eventbus.Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
	}
	p.eventBus.Publish(event)
}
