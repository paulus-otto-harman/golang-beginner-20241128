package service

import (
	"20241128/model"
	"20241128/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type shippingService struct {
	repo     repository.ShippingRepository
	validate *validator.Validate
}

type ShippingService interface {
	GetCost(shipping *model.Shipping) error
}

func NewShippingService(db *gorm.DB) ShippingService {
	return &shippingService{
		repo:     repository.NewShippingRepository(db),
		validate: validator.New(),
	}
}

func (s *shippingService) GetCost(shipping *model.Shipping) error {
	return s.repo.Cost(shipping)
}
