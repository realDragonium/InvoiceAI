package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// serve Single Page application on "/web"
	// assume static file at dist folder
	app.Static("/", "../shared_resources/")

	app.Get("/api/v1/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
