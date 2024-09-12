package di

import (
	"cheesy-events/eventbus"
	kitchenAdapter "cheesy-events/kitchen/adapter"
	kitchenDomain "cheesy-events/kitchen/domain"
	kitchenHandler "cheesy-events/kitchen/handler"
)

func ProvideKitchenHandler(eventBus *eventbus.EventBus) kitchenHandler.EventHandler {
	//sql := sqlclient.NewSqClient()
	//repository := adapter.NewPizzaRepository(sql)

	repository := kitchenAdapter.NewRepositoryDummy()
	producer := kitchenAdapter.NewProducer(eventBus)
	kitchenService := kitchenDomain.NewKitchenService(producer, &repository, 2)
	return kitchenHandler.NewEventHandler(eventBus, kitchenService) //RUN
}
