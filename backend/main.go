package main

import (
	"fmt"
	"invoiceai/database"
	"invoiceai/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := fiber.Config{}
	cfg.ErrorHandler = func(c *fiber.Ctx, err error) error {
		// TODO change later
		fmt.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{"error": "error"})
	}
	app := fiber.New(cfg)

	database.InitDatabase()

	handler.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}

}
