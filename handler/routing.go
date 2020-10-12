package handler

import (
	"github.com/gofiber/fiber/v2"
	"invoiceai/login"
)

// Setup Handler Functions here from Handler package. its bad OOP'ing but :(
func SetupRoutes(app *fiber.App) {
	app.Post("/api/v1/register", login.CreateNewUser)
	app.Get("/api/v1/auth",login.CheckValidation, login.AuthTest)
}

