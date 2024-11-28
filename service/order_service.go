package service

import (
	"20241128/model"
	"20241128/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type orderService struct {
	repo     repository.OrderRepository
	validate *validator.Validate
}

type OrderService interface {
	CreateOrder(order *model.Order) error
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{
		repo:     repository.NewOrderRepository(db),
		validate: validator.New(),
	}
}

func (s *orderService) CreateOrder(order *model.Order) error {
	if err := s.validate.Struct(order); err != nil {
		return err
	}
	return s.repo.Save(order)
}
