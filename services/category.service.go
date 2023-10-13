package services

import (
	"log"
	"synapsis_test/entities"
	"synapsis_test/inputs"

	"gorm.io/gorm"
)

type CategoryService interface {
	CreateCategory(payload *inputs.CreateProductInput) error
}

type CategoryServiceImpl struct {
	DB *gorm.DB
}

func (service *CategoryServiceImpl) CreateCategory(payload *inputs.CreateProductInput) error {
	category := entities.Category{
		Title: payload.Title,
		Desc:  payload.Description,
	}

	if err := service.DB.Create(&category).Error; err != nil {
		log.Println("Error Creating Category:", err)
		return err
	}

	return nil
}
