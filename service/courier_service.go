package service

import (
	"20241128/model"
	"20241128/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type courierService struct {
	repo     repository.CourierRepository
	validate *validator.Validate
}

type CourierService interface {
	GetAllCouriers() ([]model.Courier, error)
	CreateCourier(customer *model.Courier) error
}

func NewCourierService(db *gorm.DB) CourierService {
	return &courierService{
		repo:     repository.NewCourierRepository(db),
		validate: validator.New(),
	}
}

func (s *courierService) GetAllCouriers() ([]model.Courier, error) {
	return s.repo.FindAll()
}

func (s *courierService) CreateCourier(courier *model.Courier) error {
	if err := s.validate.Struct(courier); err != nil {
		return err
	}
	return s.repo.Save(courier)
}
