package main

import (
	"log"
	"synapsis_test/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	server.Get("/", handlers.SendHelloWorld)

	if err := server.Listen(":5000"); err != nil {
		log.Fatal(err)
	}
}
