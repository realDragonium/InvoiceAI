package handler

import (
	"invoiceai/database"
	"invoiceai/model"

	"github.com/gofiber/fiber/v2"
)

// HelloWorld is a testing api
func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Working")
}

// CreateNewUser is a function to create a new User
func CreateNewUser(c *fiber.Ctx) error {
	db := database.DBConn

	user := new(model.User)
	if err := c.BodyParser(user); err != nil {

		return c.JSON(err)

	}
	// Error Handling for Create User currenlty facing a weird Error "Unique constraint failed: users.id"
	// doesn't seem to affect the database.
	if err := db.Create(&user).Error; err != nil {
		return c.JSON(err.Error())
	}
	db.Create(&user)

	return c.JSON(fiber.Map{
		"Email":    user.Email,
		"Username": user.Username,
	})
}
