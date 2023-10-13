package routes

import (
	"synapsis_test/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAPIRoutes(server *fiber.App, handler handlers.APIHandler) {
	server.Get("/", handler.SendHelloWorld)
}
