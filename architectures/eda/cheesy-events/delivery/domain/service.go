package domain

import (
	"cheesy-events/utils/logrus"
	log "github.com/sirupsen/logrus"
)

type DeliveryStatus string

const (
	OutForDelivery DeliveryStatus = "OutForDelivery"
	Delivered      DeliveryStatus = "Delivered"
)

var logger = logrus.NewLogger("delivery")

type Delivery struct {
	ID      string
	OrderID string
	Address string
	Status  DeliveryStatus
}

type DeliveryRepository interface {
	CreateDelivery(delivery Delivery) (int64, error)
	GetDelivery(id string) (Delivery, error)
	UpdateDelivery(delivery Delivery) error
	DeleteDelivery(id string) error
}

type DeliveryProducer interface {
	SendDeliveryStarted(delivery Delivery)
	SendDeliveryCompleted(delivery Delivery)
}

type DeliveryService struct {
	repository DeliveryRepository
	producer   DeliveryProducer
}

func NewDeliveryService(repository DeliveryRepository, producer DeliveryProducer) DeliveryService {
	return DeliveryService{
		repository: repository,
		producer:   producer,
	}
}

func (ds DeliveryService) StartDelivery(delivery Delivery) error {
	if _, err := ds.repository.CreateDelivery(delivery); err != nil {
		return err
	}
	delivery.Status = OutForDelivery
	if err := ds.repository.UpdateDelivery(delivery); err != nil {
		return err
	}
	ds.producer.SendDeliveryStarted(delivery)
	logger.WithFields(log.Fields{
		"order_id": delivery.OrderID,
		"status":   delivery.Status,
	}).Info("Delivering pizza....")

	delivery.Status = Delivered
	err := ds.repository.UpdateDelivery(delivery)
	if err != nil {
		return err
	}
	ds.producer.SendDeliveryCompleted(delivery)
	logger.WithFields(log.Fields{
		"order_id": delivery.OrderID,
		"status":   delivery.Status,
	}).Info("PizzaName delivered")
	return nil
}
