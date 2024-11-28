package database

import (
	"20241128/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		seeds := dataSeeds()
		for _, seed := range seeds {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seed).Error
			if err != nil {
				name := reflect.TypeOf(seed).String()
				errorMessage := err.Error()
				return fmt.Errorf("%s seeder fail with %s", name, errorMessage)
			}
		}
		return nil
	})

}

func dataSeeds() []interface{} {
	return []interface{}{
		model.CourierSeed(),
	}
}
