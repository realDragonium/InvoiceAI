package router

import (
	"invoiceai/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//stolen from https://github.com/gofiber/recipes/tree/master/auth-jwt
// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	//auth := api.Group("/auth")
	//auth.Post("/login", handler.Login)

	// User
	//user := api.Group("/user")
	//user.Get("/:id", handler.GetUser)
	//user.Post("/", handler.CreateUser)
	//user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	//user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Product
	//product := api.Group("/product")
	//product.Get("/", handler.GetAllProducts)
	//product.Get("/:id", handler.GetProduct)
	//product.Post("/", middleware.Protected(), handler.CreateProduct)
	//product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)

	app.Static("/", "./frontend/public")

	app.Get("/web/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./frontend/public/index.html")
	})

	app.Get("/api/v1/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})
}
