package handler

import (
	"github.com/gofiber/fiber/v2"
	"invoiceai/auth"
	"invoiceai/customer"
)

// Setup Handler Functions here from Handler package. its bad OOP'ing but :(
func SetupRoutes(app *fiber.App) {

	// Proxy for front end working now.
	//app.Get("/", proxy.Forward("http://localhost:5000/"))

	// Routes
	app.Post("/api/v1/register", auth.CreateNewUser)
	app.Post("/api/v1/login", auth.UserLogin)
	app.Get("/api/v1/auth", auth.CheckValidation, auth.TokenUserIDValidation)

	app.Post("api/customer/create", customer.Create)
	app.Get("api/customer/all", customer.GetAllCustomer)

	// 404 Handler this route needs to be the last one otherwise it will match all routes.
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"error": "page not found"}) // => 404 "Not Found"
	})

}
