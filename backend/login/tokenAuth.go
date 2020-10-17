package login

import (
	"fmt"
	"github.com/brianvoe/sjwt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)




// Middleware to check if JWT is valid
func CheckValidation(c *fiber.Ctx) error{
	token:=c.Cookies("token")
	verified:= verifyToken(token)
	if verified{
		parsedData := parseToken(token)
		convertTokenToUint,_ := strconv.ParseUint(parsedData, 10, 32)
		uintTokenId:= uint(convertTokenToUint)
		fmt.Println(uintTokenId)
		token := generateJWT(uintTokenId)
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = token
		cookie.HTTPOnly = true
		cookie.Expires = time.Now().Add(time.Minute * 120)
		fmt.Println(token)
		c.Cookie(cookie)

		//c.Redirect("/api/v1/register")

		return c.Next()
	}else{
		return c.SendString("Please Log In")
	}
}



// Functionality for generating JWT tokens parsing them and verifying them
func generateJWT(userID uint) (token string){
	claims := sjwt.New()
	claims.Set("ID", userID)
	claims.SetExpiresAt(time.Now().Add(time.Minute * 120))

	// Generate jwt
	secretKey := []byte("122334")
	jwt := claims.Generate(secretKey)
	return jwt
}

func parseToken(token string) (parsedData string){
	// Parse jwt
	jwt := token
	claims, _ := sjwt.Parse(jwt)

	// Get claims
	parsedData, err := claims.GetStr("ID") // John Doe
	if err != nil{
		fmt.Print("something went wrong")
	}
	return parsedData
}

func verifyToken(token string) (verified bool){
	secretKey := []byte("122334")

	// Verify that the secret signature is valid
	hasVerified := sjwt.Verify(token, secretKey)

	// Parse jwt
	claims, _ := sjwt.Parse(token)

	// Validate will check(if set) Expiration At and Not Before At dates
	err := claims.Validate()
	if err != nil{
		fmt.Print("Error with validating token")
	}
	return hasVerified
}