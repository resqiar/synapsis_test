package handlers

import (
	"synapsis_test/inputs"
	"synapsis_test/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler interface {
	CreateCategory(c *fiber.Ctx) error
}

type CategoryHandlerImpl struct {
	UtilService     services.UtilService
	CategoryService services.CategoryService
}

func (handler *CategoryHandlerImpl) CreateProduct(c *fiber.Ctx) error {
	var payload inputs.CreateProductInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// validate the payload using class-validator package
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).SendString(err)
	}

	if err := handler.CategoryService.CreateCategory(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
