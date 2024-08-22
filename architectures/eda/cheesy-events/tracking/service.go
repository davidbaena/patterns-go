package tracking

import (
	"cheesy-events/delivery"
	"cheesy-events/eventbus"
	"cheesy-events/kitchen"
	"cheesy-events/orders"
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
)

var logger = logrus.NewLogger("tracking")

// OrderTrackingService represents the service responsible for tracking order status
type OrderTrackingService struct {
	eventBus *eventbus.EventBus
}

// NewOrderTrackingService creates a new instance of the order tracking service
func NewOrderTrackingService(eventBus *eventbus.EventBus) *OrderTrackingService {

	// Create channels for the events
	orderPlacedChan := make(chan eventbus.Event)
	orderAcceptedChan := make(chan eventbus.Event)
	pizzaBeingPreparedChan := make(chan eventbus.Event)
	outForDeliveryChan := make(chan eventbus.Event)
	deliveredChan := make(chan eventbus.Event)

	// subscribe the channels to the events
	eventBus.Subscribe("OrderPlaced", orderPlacedChan)
	eventBus.Subscribe("OrderAccepted", orderAcceptedChan)
	eventBus.Subscribe("PizzaBeingPrepared", pizzaBeingPreparedChan)
	eventBus.Subscribe("OutForDelivery", outForDeliveryChan)
	eventBus.Subscribe("Delivered", deliveredChan)

	ots := &OrderTrackingService{
		eventBus: eventBus,
	}

	go ots.subscribe(orderPlacedChan)
	go ots.subscribe(orderAcceptedChan)
	go ots.subscribe(pizzaBeingPreparedChan)
	go ots.subscribe(outForDeliveryChan)
	go ots.subscribe(deliveredChan)

	return ots
}

func (ots *OrderTrackingService) onOrderPlaced(orderPlacedEvent orders.OrderPlacedEvent) {
	logger.WithFields(log.Fields{
		"order_id": orderPlacedEvent.OrderID,
		"Status":   orderPlacedEvent.Status,
	}).Info("Order placed")
}

func (ots *OrderTrackingService) onOrderAccepted(orderAcceptedEvent orders.OrderAcceptedEvent) {
	logger.WithFields(log.Fields{
		"order_id": orderAcceptedEvent.OrderID,
		"Status":   orderAcceptedEvent.Status,
	}).Info("Order accepted")
}

func (ots *OrderTrackingService) onPizzaBeingPrepared(pizzaBeingPreparedEvent kitchen.PizzaBeingPreparedEvent) {
	logger.WithFields(log.Fields{
		"order_id": pizzaBeingPreparedEvent.OrderID,
		"Status":   pizzaBeingPreparedEvent.Status,
	}).Info("Pizza being prepared")
}

func (ots *OrderTrackingService) onOutForDelivery(outForDeliveryEvent delivery.OutForDeliveryEvent) {
	logger.WithFields(log.Fields{
		"order_id": outForDeliveryEvent.OrderID,
		"Status":   outForDeliveryEvent.Status,
	}).Info("Out for delivery")
}

func (ots *OrderTrackingService) onDelivered(deliveredEvent delivery.DeliveredEvent) {
	logger.WithFields(log.Fields{
		"order_id": deliveredEvent.OrderID,
		"Status":   deliveredEvent.Status,
	}).Info("Delivered")
}

// subscribe listens for order status events and updates the order status
func (ots *OrderTrackingService) subscribe(eventChan <-chan eventbus.Event) {
	for event := range eventChan {
		switch e := event.Data.(type) {
		case orders.OrderPlacedEvent:
			ots.onOrderPlaced(e)
		case orders.OrderAcceptedEvent:
			ots.onOrderAccepted(e)
		case kitchen.PizzaBeingPreparedEvent:
			ots.onPizzaBeingPrepared(e)
		case delivery.OutForDeliveryEvent:
			ots.onOutForDelivery(e)
		case delivery.DeliveredEvent:
			ots.onDelivered(e)
		default:
			logger.Error("Unknown event type")
		}
	}
}
