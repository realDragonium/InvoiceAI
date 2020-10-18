package main

import (
	"fmt"
	"invoiceai/database"
	"invoiceai/handler"
	"invoiceai/login"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Still need this for front end
	app.Static("/", "../frontend/public")
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins: "*",
	//	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept",
	//	AllowCredentials: true,
	//	ExposeHeaders: "Set-Cookie",
	//
	//}))
	database.InitDatabase()
	login.Migrate()
	handler.SetupRoutes(app)
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
