package auth

import (
	"fmt"
	"github.com/brianvoe/sjwt"
	"github.com/gofiber/fiber/v2"
	"time"
)

var secretKey = []byte("122334")

// Middleware to check if JWT is valid
func CheckValidation(c *fiber.Ctx) error {
	token := c.Cookies("token")
	// Verify that the secret signature is valid
	hasVerified := sjwt.Verify(token, secretKey)
	claims, _ := sjwt.Parse(token)

	// Validate will check(if set) Expiration At and Not Before At dates
	err := claims.Validate()
	if err != nil {
		return err
	}

	if !hasVerified {
		return c.Status(401).JSON(fiber.Map{"error": "Please Log In"})
	}
	userID := userIDFromCtx(c)
	fmt.Println(userID)
	sendJWT(c, userID)
	return c.Next()
}

// Generates token and sets cookie with it for 2 hours when a user logs in
func sendJWT(c *fiber.Ctx, userID string) {
	claims := sjwt.New()
	claims.Set("ID", userID)
	claims.SetExpiresAt(time.Now().Add(time.Minute * 120))
	token := claims.Generate(secretKey)

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(time.Minute * 120)
	c.Cookie(cookie)
}

func userIDFromCtx(c *fiber.Ctx) string {
	token := c.Cookies("token")
	claims, _ := sjwt.Parse(token)

	id, err := claims.GetStr("ID")
	if err != nil {
		fmt.Print("something went wrong")
	}
	return id
}
