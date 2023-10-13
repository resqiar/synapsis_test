package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Title       string
	Description string
	ImageURL    string

	Categories []Category `gorm:"many2many:product_categories;"`
}
