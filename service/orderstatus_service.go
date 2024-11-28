package service

import (
	"20241128/model"
	"20241128/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type orderStatusService struct {
	repo     repository.OrderStatusRepository
	validate *validator.Validate
}

type OrderStatusService interface {
	CreateOrderStatus(orderStatus *model.OrderStatus) error
	GetByOrderId(orderId uint) ([]model.OrderStatus, error)
}

func NewOrderStatusService(db *gorm.DB) OrderStatusService {
	return &orderStatusService{
		repo:     repository.NewOrderStatusRepository(db),
		validate: validator.New(),
	}
}

func (s *orderStatusService) CreateOrderStatus(orderStatus *model.OrderStatus) error {
	if err := s.validate.Struct(orderStatus); err != nil {
		return err
	}
	return s.repo.Save(orderStatus)
}

func (s *orderStatusService) GetByOrderId(orderId uint) ([]model.OrderStatus, error) {
	return s.repo.GetByOrderId(orderId)
}
