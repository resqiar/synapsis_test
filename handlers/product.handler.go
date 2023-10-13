package handlers

import (
	"synapsis_test/inputs"
	"synapsis_test/services"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	CreateProduct(c *fiber.Ctx) error
}

type ProductHandlerImpl struct {
	UtilService    services.UtilService
	ProductService services.ProductService
}

func (handler *ProductHandlerImpl) CreateProduct(c *fiber.Ctx) error {
	var payload inputs.CreateProductInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// validate the payload using class-validator package
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).SendString(err)
	}

	if err := handler.ProductService.CreateProduct(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
