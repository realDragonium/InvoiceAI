package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"invoiceai/database"
	"invoiceai/handler"
	"invoiceai/login"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Default config
	app.Use(cors.New())
	database.InitDatabase()
	login.Migrate()
	handler.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
