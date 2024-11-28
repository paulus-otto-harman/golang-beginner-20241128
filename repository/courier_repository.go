package repository

import (
	"20241128/model"
	"gorm.io/gorm"
	"log"
)

type CourierRepository interface {
	FindAll() ([]model.Courier, error)
	Save(customer *model.Courier) error
}

type courierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(db *gorm.DB) CourierRepository {
	return &courierRepository{db: db}
}

func (r *courierRepository) FindAll() ([]model.Courier, error) {
	var couriers []model.Courier
	err := r.db.Find(&couriers).Error
	log.Println(couriers)
	return couriers, err
}

func (r *courierRepository) Save(courier *model.Courier) error {
	return r.db.Create(courier).Error
}
