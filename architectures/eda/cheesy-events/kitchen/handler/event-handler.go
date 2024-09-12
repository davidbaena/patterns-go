package handler

import (
	"cheesy-events/eventbus"
	"cheesy-events/kitchen/domain"
	"cheesy-events/orders"
	"cheesy-events/utils/logrus"
)

var logger = logrus.NewLogger("KitchenEventHandler")

type EventHandler struct {
	eventBus       *eventbus.EventBus
	kitchenService domain.KitchenService
}

func NewEventHandler(eventBus *eventbus.EventBus, kitchenService domain.KitchenService) *EventHandler {
	handler := &EventHandler{
		eventBus:       eventBus,
		kitchenService: kitchenService,
	}
	orderAcceptedChan := make(chan eventbus.Event)
	eventBus.Subscribe("OrderAccepted", orderAcceptedChan)
	go handler.subscribe(orderAcceptedChan)
	return handler
}

func (h *EventHandler) subscribe(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		orderAcceptedEvent, ok := event.Data.(orders.OrderAcceptedEvent)
		if !ok {
			logger.Error("Invalid event data")
			continue
		}

		pizza := domain.Pizza{
			OrderId: orderAcceptedEvent.OrderID,
			Name:    orderAcceptedEvent.Pizza,
			Status:  domain.OrderReceived,
		}
		h.kitchenService.PreparePizza(pizza)
	}
}
