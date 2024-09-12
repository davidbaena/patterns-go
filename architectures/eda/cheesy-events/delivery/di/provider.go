package di

import (
	"cheesy-events/delivery/adapter"
	"cheesy-events/delivery/domain"
	"cheesy-events/delivery/handler"
	"cheesy-events/eventbus"
)

func ProvideDeliveryHandler(eventBus *eventbus.EventBus) handler.DeliveryEventHandler {
	repository := adapter.NewDeliveryRepositoryDummy()
	producer := adapter.NewDeliveryProducer(eventBus)
	service := domain.NewDeliveryService(&repository, producer)
	return handler.NewDeliveryEventHandler(eventBus, service)
}
