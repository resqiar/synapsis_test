package services

import (
	"log"
	"synapsis_test/entities"
	"synapsis_test/inputs"

	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(payload *inputs.CreateProductInput) error
}

type ProductServiceImpl struct {
	DB *gorm.DB
}

func (service *ProductServiceImpl) CreateProduct(payload *inputs.CreateProductInput) error {
	product := entities.Product{
		Title:       payload.Title,
		Description: payload.Description,
		ImageURL:    payload.ImageURL,
	}

	if err := service.DB.Create(&product).Error; err != nil {
		log.Println("Error Creating Product:", err)
		return err
	}

	return nil
}
