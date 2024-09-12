package handler

import (
	"cheesy-events/delivery/domain"
	"cheesy-events/eventbus"
	"cheesy-events/kitchen/adapter"
	"cheesy-events/utils/logrus"
)

var logger = logrus.NewLogger("DeliveryEventHandler")

type DeliveryEventHandler struct {
	eventBus        *eventbus.EventBus
	deliveryService domain.DeliveryService
}

func NewDeliveryEventHandler(eventBus *eventbus.EventBus, deliveryService domain.DeliveryService) DeliveryEventHandler {
	handler := DeliveryEventHandler{
		eventBus:        eventBus,
		deliveryService: deliveryService,
	}
	pizzaPrepared := make(chan eventbus.Event)
	eventBus.Subscribe("PizzaPreparedEvent", pizzaPrepared)
	go handler.subscribe(pizzaPrepared)
	return handler
}

func (h *DeliveryEventHandler) subscribe(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		pizzaPreparedEvent, ok := event.Data.(adapter.PizzaPreparedEvent)
		if !ok {
			logger.Error("Invalid event data")
			continue
		}

		delivery := domain.Delivery{
			OrderID: pizzaPreparedEvent.OrderID,
			Status:  domain.OutForDelivery,
		}
		err := h.deliveryService.StartDelivery(delivery)
		if err != nil {
			logger.Error("Failed to start delivery: ", err)
		}
	}
}
