package repository

import (
	"20241128/model"
	"gorm.io/gorm"
	"log"
)

type OrderRepository interface {
	Save(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Save(order *model.Order) error {
	firstDeliveryStatus := "Menunggu Pickup"
	initialLocation := "Rumah Penjual"
	order.Status = firstDeliveryStatus
	//if err := r.db.Preload("Courier").Find(&order).Error; err != nil {
	//	log.Fatalf("failed to fetch order: %v", err)
	//}
	order.OrderStatus = append(order.OrderStatus, model.OrderStatus{Status: firstDeliveryStatus, Location: initialLocation})
	q := r.db.Create(order)

	if err := r.db.Preload("Courier").Find(&order).Error; err != nil {
		log.Fatalf("failed to fetch order: %v", err)
	}
	return q.Error
}
