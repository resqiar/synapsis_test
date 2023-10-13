package handlers

import (
	"synapsis_test/inputs"
	"synapsis_test/services"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	SendRegister(c *fiber.Ctx) error
}

type AuthHandlerImpl struct {
	AuthService services.AuthService
	UtilService services.UtilService
}

func (handler *AuthHandlerImpl) SendRegister(c *fiber.Ctx) error {
	// body payload
	var payload inputs.CreateUserInput

	// bind the body into payload
	if err := c.BodyParser(&payload); err != nil {
		// send raw error (unprocessable entity)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// validate the payload using class-validator package
	if err := handler.UtilService.ValidateInput(payload); err != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err,
		})
	}

	if err := handler.AuthService.Register(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
