package repository

import (
	"20241128/model"
	"gorm.io/gorm"
)

type OrderStatusRepository interface {
	Save(order *model.OrderStatus) error
	GetByOrderId(orderId uint) ([]model.OrderStatus, error)
}

type orderStatusRepository struct {
	db *gorm.DB
}

func NewOrderStatusRepository(db *gorm.DB) OrderStatusRepository {
	return &orderStatusRepository{db: db}
}

func (r *orderStatusRepository) Save(orderStatus *model.OrderStatus) error {
	return r.db.Save(orderStatus).Error
}

func (r *orderStatusRepository) GetByOrderId(orderId uint) ([]model.OrderStatus, error) {
	var orderStatus []model.OrderStatus
	err := r.db.Find(&orderStatus, "order_id=?", orderId).Error
	return orderStatus, err
}
