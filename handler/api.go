package handler

import (
	"github.com/gofiber/fiber/v2"
)

// HelloWorld is a testing api
func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Working")
}
