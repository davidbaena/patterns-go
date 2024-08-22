package delivery

import (
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
	"time"

	"cheesy-events/eventbus"
	"cheesy-events/kitchen"
)

var logger = logrus.NewLogger("kitchen")

// OutForDeliveryEvent represents the event structure for out for delivery status
type OutForDeliveryEvent struct {
	OrderID string
	Status  string
}

// DeliveredEvent represents the event structure for delivered status
type DeliveredEvent struct {
	OrderID string
	Status  string
}

// DeliveryService represents the service responsible for pizza delivery
type DeliveryService struct {
	eventBus *eventbus.EventBus
}

// NewDeliveryService creates a new instance of the delivery service
func NewDeliveryService(eventBus *eventbus.EventBus) *DeliveryService {
	pizzaPreparedChan := make(chan eventbus.Event)
	eventBus.Subscribe("PizzaPrepared", pizzaPreparedChan)

	ds := &DeliveryService{
		eventBus: eventBus,
	}

	go ds.subscribe(pizzaPreparedChan)

	return ds
}

func (ds *DeliveryService) onPizzaPrepared(pizzaPreparedEvent kitchen.PizzaPreparedEvent) {
	// Create the out for delivery event
	outForDeliveryEvent := eventbus.Event{
		Type:      "OutForDelivery",
		Timestamp: time.Now(),
		Data: OutForDeliveryEvent{
			OrderID: pizzaPreparedEvent.OrderID,
			Status:  "Out For Delivery",
		},
	}
	ds.eventBus.Publish(outForDeliveryEvent)
	logger.WithFields(log.Fields{
		"order_id": pizzaPreparedEvent.OrderID,
		"pizza":    pizzaPreparedEvent.Pizza,
	}).Info("Delivering pizza....")

	time.Sleep(1 * time.Second)

	// Create the delivered event
	deliveredEvent := eventbus.Event{
		Type:      "Delivered",
		Timestamp: time.Now(),
		Data: DeliveredEvent{
			OrderID: pizzaPreparedEvent.OrderID,
			Status:  "Delivered",
		},
	}
	logger.WithFields(log.Fields{
		"order_id": pizzaPreparedEvent.OrderID,
		"pizza":    pizzaPreparedEvent.Pizza,
	}).Info("Pizza delivered")
	ds.eventBus.Publish(deliveredEvent)
}

// subscribe delivers a pizza when a pizza prepared event is received
func (ds *DeliveryService) subscribe(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		pizzaPreparedEvent, ok := event.Data.(kitchen.PizzaPreparedEvent)
		if !ok {
			logger.Error("Invalid event data")
			continue
		}
		ds.onPizzaPrepared(pizzaPreparedEvent)
	}
}
