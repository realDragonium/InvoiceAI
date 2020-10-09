package main

import (
	"invoiceai/database"
	"invoiceai/handler"

	"github.com/gofiber/fiber/v2"
)

// Setup Handler Functions here from Handler package. its bad OOP'ing but :(
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/test", handler.HelloWorld)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	database.InitDatabase()
	app.Listen(":3000")
}
