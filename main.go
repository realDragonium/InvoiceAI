package main

import (
	"invoiceai/database"
	"invoiceai/handler"
	"invoiceai/login"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDatabase()
	login.Migrate()
	handler.SetupRoutes(app)
	app.Listen(":3000")
}
