package main

import (
	"log"
	"synapsis_test/db"
	"synapsis_test/handlers"
	"synapsis_test/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	// Init DB
	_, err := db.InitDatabase()
	if err != nil {
		log.Fatal("DATABASE FAILED: ", err)
	}

	// Init handlers
	APIHandler := handlers.APIHandlerImpl{}
	AuthHandler := handlers.AuthHandlerImpl{}

	// Init routes
	routes.InitAPIRoutes(server, &APIHandler)
	routes.InitAuthRoutes(server, &AuthHandler)

	if err := server.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
