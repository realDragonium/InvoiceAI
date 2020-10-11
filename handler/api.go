package handler

import (
	"invoiceai/database"
	"invoiceai/model"
	"strings"

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
	var users []model.User

	// since Gorm Docs doesn't seem to have a way to make emails Unique i wrote this
	// to check the database if the email exists if it does it will give an error.
	var checkIfEmailExist model.User
	db.First(&checkIfEmailExist, "email = ?", user.Email)
	if strings.Contains(user.Email, checkIfEmailExist.Email) {
		return c.JSON(fiber.Map{
			"Error": "Email Already in Use",
		})
	}
	db.Create(&user)

	return c.JSON(users)
}
