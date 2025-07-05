package delivery

import "gorm.io/gorm"

type DeliveryService struct {
	Db      *gorm.DB
	Request *Request
}

func NewDeliveryService(db *gorm.DB, request *Request) *DeliveryService {
	return &DeliveryService{Db: db, Request: request}
}

func (service *DeliveryService) GetDelivery() (interface{}, error) {
	deliveryRepository := NewDeliveryRepository(service.Db, service.Request)
	return deliveryRepository.GetDelivery()
}
