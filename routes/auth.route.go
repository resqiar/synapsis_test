package routes

import (
	"synapsis_test/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitAuthRoutes(server *fiber.App, handler handlers.AuthHandler) {
	auth := server.Group("auth")

	auth.Post("/register", handler.SendRegister)
}
