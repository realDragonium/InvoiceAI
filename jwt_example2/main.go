package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

const jwtsecret = "asecret"

func authRequire() func(ctx *fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) {
			ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(jwtsecret),
	})
}

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello")
	})

	app.Post("/login", login)

	app.Get("/hello", authRequire(), func(ctx *fiber.Ctx) {
		user := ctx.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["sub"].(string)

		ctx.Send(fmt.Sprintf("hello user with id: %v", id))
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func login(ctx *fiber.Ctx) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}
	if body.Email != "bob@gmail.com" || body.Password != "password1234" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7)

	s, err := token.SignedString([]byte(jwtsecret))
	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
		return
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}{
			Id:    1,
			Email: "bob@gmail.com",
		},
	})

}
