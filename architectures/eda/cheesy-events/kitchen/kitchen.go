package kitchen

import (
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"

	"cheesy-events/eventbus"
	"cheesy-events/orders"
)

var logger = logrus.NewLogger("kitchen")

// PizzaPreparedEvent represents the event structure for pizza preparation
type PizzaPreparedEvent struct {
	OrderID string
	Pizza   string
}

// PizzaBeingPreparedEvent represents the event structure for pizza preparation status
type PizzaBeingPreparedEvent struct {
	OrderID string
	Status  string
}

// KitchenService represents the service responsible for pizza preparation
type KitchenService struct {
	eventBus       *eventbus.EventBus
	pizzasPrepared int
	maxPizzas      int
}

// NewKitchenService creates a new instance of the kitchen service
func NewKitchenService(eventBus *eventbus.EventBus, maxPizzas int) *KitchenService {
	orderAcceptedChan := make(chan eventbus.Event)
	eventBus.Subscribe("OrderAccepted", orderAcceptedChan)

	ks := &KitchenService{
		eventBus:  eventBus,
		maxPizzas: maxPizzas,
	}
	go ks.subscribe(orderAcceptedChan)

	return ks
}

// subscribe prepares a pizza and publishes a pizza prepared event
func (ks *KitchenService) subscribe(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		orderAcceptedEvent, ok := event.Data.(orders.OrderAcceptedEvent)
		if !ok {
			logger.Error("Invalid event data")
			continue
		}
		ks.onOrderAccepted(orderAcceptedEvent)
	}
}

func (ks *KitchenService) onOrderAccepted(orderAcceptedEvent orders.OrderAcceptedEvent) {

	// Check if maximum number of pizzas have been prepared
	if ks.pizzasPrepared >= ks.maxPizzas {
		logger.Error("Maximum number of pizzas prepared. Cannot prepare more.")
		return
	}

	// Create the pizza being prepared event
	preparingEvent := eventbus.Event{
		Type:      "PizzaBeingPrepared",
		Timestamp: time.Now(),
		Data: PizzaBeingPreparedEvent{
			OrderID: orderAcceptedEvent.OrderID,
			Status:  "Pizza Being Prepared",
		},
	}
	ks.eventBus.Publish(preparingEvent)

	logger.WithFields(log.Fields{
		"order_id": orderAcceptedEvent.OrderID,
		"pizza":    orderAcceptedEvent.Pizza,
	}).Info("Preparing pizza...")

	time.Sleep(1 * time.Second)
	// Create the pizza prepared event
	event := eventbus.Event{
		Type:      "PizzaPrepared",
		Timestamp: time.Now(),
		Data: PizzaPreparedEvent{
			OrderID: orderAcceptedEvent.OrderID,
			Pizza:   orderAcceptedEvent.Pizza,
		},
	}
	logger.WithFields(log.Fields{
		"order_id": orderAcceptedEvent.OrderID,
		"pizza":    orderAcceptedEvent.Pizza,
	}).Info("Pizza prepared")
	ks.eventBus.Publish(event)

	// Increment the number of pizzas prepared
	ks.pizzasPrepared++
}
