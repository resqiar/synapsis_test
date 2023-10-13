package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	SendRegister(c *fiber.Ctx) error
}

type AuthHandlerImpl struct{}

func (handler *AuthHandlerImpl) SendRegister(c *fiber.Ctx) error {
	return c.SendString("Hello from auth")
}
