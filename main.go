package main

import (
	"github.com/gofiber/fiber/v2"
	"invoiceai/database"
	"invoiceai/router"
	"log"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))

}
