package routes

import (
	"synapsis_test/handlers"

	"github.com/gofiber/fiber/v2"
)

func InitCategoryRoutes(server *fiber.App, handler handlers.ProductHandler) {
	product := server.Group("category")
	product.Post("/create", handler.CreateProduct)
}
