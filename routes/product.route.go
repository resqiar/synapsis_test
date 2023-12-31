package routes

import (
	"synapsis_test/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitProductRoutes(server *fiber.App, handler handlers.ProductHandler) {
	product := server.Group("product")
	product.Post("/create", handler.CreateProduct)
}
