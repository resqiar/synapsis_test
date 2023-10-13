package main

import (
	"log"
	"synapsis_test/db"
	"synapsis_test/handlers"
	"synapsis_test/routes"
	"synapsis_test/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	server := fiber.New()

	// Init DB
	db, err := db.InitDatabase()
	if err != nil {
		log.Fatal("DATABASE FAILED: ", err)
	}

	// Init all modules (services, handlers, routes)
	initModules(db, server)

	if err := server.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}

func initModules(db *gorm.DB, server *fiber.App) {
	// services
	utilService := services.UtilServiceImpl{}
	authService := services.AuthServiceImpl{
		DB:          db,
		UtilService: &utilService,
	}
	productService := services.ProductServiceImpl{
		DB: db,
	}

	// handlers
	productHandler := handlers.ProductHandlerImpl{
		UtilService:    &utilService,
		ProductService: &productService,
	}
	authHandler := handlers.AuthHandlerImpl{
		UtilService: &utilService,
		AuthService: &authService,
	}

	// routes
	routes.InitProductRoutes(server, &productHandler)
	routes.InitAuthRoutes(server, &authHandler)
}
