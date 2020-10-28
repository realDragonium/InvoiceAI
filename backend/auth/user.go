package auth

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"invoiceai/database"
)

// Database table
type user struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	ID       string `json:"id"`
	Verified bool   `json:"verified"`
	Password string `json:"password"`
}

type registerUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateNewUser is a function to create a new User
func CreateNewUser(c *fiber.Ctx) error {
	user := new(registerUser)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	if _, err := database.DB.Exec(context.Background(), `insert into "user" (email, name, password) values ($1, $2, $3)`, user.Email, user.Name, hashedPassword); err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"success": "success"})
}

// UserLogin functionality to login a user and generate a JWT for them
func UserLogin(c *fiber.Ctx) error {
	user := new(loginUser)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	var hashedPassword string
	if err := database.DB.QueryRow(context.Background(), `select password from public.user where email = $1;`, user.Email).Scan(&hashedPassword); err != nil {
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return c.Status(404).JSON(fiber.Map{"error": "not allowed"})
	}

	var userID string
	if err := database.DB.QueryRow(context.Background(), `select id from public.user where email = $1;`, user.Email).Scan(&userID); err != nil {
		return err
	}

	sendJWT(c, userID)
	return c.Status(201).JSON(fiber.Map{"success": "success"})
}

// verifying if the userID in JWT is valid
func TokenUserIDValidation(c *fiber.Ctx) error {
	var exists int
	userID := userIDFromCtx(c)
	if err := database.DB.QueryRow(context.Background(), `select count(id) from public.user where id = $1;`, userID).Scan(&exists); err != nil {
		return err
	}
	if exists == 0 {
		return c.Status(401).JSON(fiber.Map{"error": "Please Log In"})
	}
	return c.Status(200).JSON(fiber.Map{"success": "Logged In"})
}

// Will be used later by changing passwords
// HashPassword is a function to hash a Password will be used in auth system  using example from  https://gowebexamples.com/password-hashing/
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
