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
	user := entities.Product{
		Title:       payload.Title,
		Description: payload.Description,
		ImageURL:    payload.ImageURL,
	}

	if err := service.DB.Create(&user).Error; err != nil {
		log.Println("Error Creating Product:", err)
		return err
	}

	return nil
}
