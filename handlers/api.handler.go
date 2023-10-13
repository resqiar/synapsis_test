package handlers

import "github.com/gofiber/fiber/v2"

type APIHandler interface {
	SendHelloWorld(c *fiber.Ctx) error
}

type APIHandlerImpl struct{}

func (handler *APIHandlerImpl) SendHelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
