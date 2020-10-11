package handler

import (
	"github.com/gofiber/fiber/v2"
	"invoiceai/login"
)

// Setup Handler Functions here from Handler package. its bad OOP'ing but :(
func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/test", helloWorld)
	app.Post("/api/v1/user", login.CreateNewUser)
}

// HelloWorld is a testing api
func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Working")
}
