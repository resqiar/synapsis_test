package handlers

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	SendHelloWorld(c *fiber.Ctx) error
}

type ProductHandlerImpl struct{}

func (handler *ProductHandlerImpl) SendHelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
