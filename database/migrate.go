package database

import (
	"20241128/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Courier{}, &model.Order{}, &model.OrderStatus{})
	return err
}
