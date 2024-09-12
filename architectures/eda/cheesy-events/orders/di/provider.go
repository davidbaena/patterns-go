package di

import (
	"cheesy-events/eventbus"
	"cheesy-events/orders/adapter"
	"cheesy-events/orders/domain"
	"cheesy-events/orders/handler"
)

func ProvideOrderHandler(eventBus *eventbus.EventBus) handler.OrderEventHandler {
	producer := adapter.NewOrderProducer(eventBus)
	service := domain.NewOrderService(eventBus, producer)
	return handler.NewOrderEventHandler(eventBus, service)
}
