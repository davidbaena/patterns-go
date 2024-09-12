package handler

import (
	"cheesy-events/eventbus"
	"cheesy-events/orders/domain"
	"cheesy-events/utils/logrus"
)

var logger = logrus.NewLogger("OrderEventHandler")

// OrderPlacedEvent represents the event structure for order placement
type OrderPlacedEvent struct {
	OrderID string
	Pizza   string
}

type OrderEventHandler struct {
	eventBus     *eventbus.EventBus
	orderService domain.OrderService
}

func NewOrderEventHandler(eventBus *eventbus.EventBus, orderService domain.OrderService) OrderEventHandler {
	handler := OrderEventHandler{
		eventBus:     eventBus,
		orderService: orderService,
	}
	orderPlacedChan := make(chan eventbus.Event)
	eventBus.Subscribe("OrderPlaced", orderPlacedChan)
	go handler.subscribe(orderPlacedChan)
	return handler
}

func (h *OrderEventHandler) subscribe(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		orderPlacedEvent, ok := event.Data.(OrderPlacedEvent)
		if !ok {
			logger.Error("Invalid event data")
			continue
		}

		order := domain.Order{
			ID:    orderPlacedEvent.OrderID,
			Pizza: orderPlacedEvent.Pizza,
		}

		h.orderService.PlaceOrder(order)
	}
}
