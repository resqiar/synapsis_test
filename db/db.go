package db

import (
	"synapsis_test/entities"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/data.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrations - should be off in prod
	db.AutoMigrate(&entities.User{}, &entities.Product{}, &entities.Category{}, &entities.Cart{}, &entities.CartItem{})

	return db, nil
}
