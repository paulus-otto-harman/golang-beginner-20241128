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
	if err := r.db.Preload("Courier").Table("couriers").First(&order, "id = ?", order.CourierID).Error; err != nil {
		log.Fatalf("failed to fetch order: %v", err)
	}

	order.Status = firstDeliveryStatus
	order.OrderStatus = append(order.OrderStatus, model.OrderStatus{Status: firstDeliveryStatus, Location: initialLocation})

	return r.db.Create(order).Error
}
