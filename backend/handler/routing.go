package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"invoiceai/auth"
)

// Setup Handler Functions here from Handler package. its bad OOP'ing but :(
func SetupRoutes(app *fiber.App) {

	// Proxy for front end
	app.Get("/test", proxy.Forward("http://127.0.0.1:5000/"))

	app.Post("/api/v1/register", auth.CreateNewUser)
	app.Post("/api/v1/login", auth.UserLogin)
	app.Get("/api/v1/auth", auth.CheckValidation, auth.AuthTest)

	// 404 Handler this route needs to be the last one otherwise it will match all routes.
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"error": "page not found", "status": 404}) // => 404 "Not Found"
	})

}
