package login

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"invoiceai/database"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"time"
)

func Migrate() {
	database.DBConn.AutoMigrate(&user{})
}

// User model for auth system in the future?
type user struct {
	gorm.Model
	Username string `json:"username" `
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}


// Routes
// CreateNewUser is a function to create a new User
func CreateNewUser(c *fiber.Ctx) error {
	db := database.DBConn

	user := new(user)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(err)
	}

	// Error Handling for Create User currenlty facing a weird Error "Unique constraint failed: users.id"
	// doesn't seem to affect the database.
	if err := db.Create(&user).Error; err != nil {
		//TODO a better way for this
		return c.JSON(err.Error())
	}

	// Generates token and sets cookie with it for 2 hours when a user signs up then we can redirect them
	token := generateJWT(user.ID)
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(time.Minute * 120)
	c.Cookie(cookie)
	return c.JSON(fiber.Map{
		"email":    user.Email,
		"username": user.Username,
		"token": token,
	})


}



// Just a testing route to verify Auth is working
func AuthTest(c *fiber.Ctx) error{
	return c.JSON("Logged In!")
}




// BeforeCreate is a function to add password hashing and User validation see https://gorm.io/docs/hooks.html
func (u *user) BeforeCreate(tx *gorm.DB) (err error) {

	// Hasing of Password before saving to database
	hashedPassword, err := hashPassword(u.Password)
	if err != nil {
		return errors.New("something went wrong")
	}

	// Create User validation for password
	if len(u.Password) < 8 {
		return errors.New("Password has to be longer than 8 Characters")
	}

	// Create User validation for Username
	if len(u.Username) < 8 {
		return errors.New("Username has to be longer than 8 Characters")
	}

	// Create User validation for Email
	re := regexp.MustCompile("^\\S+@\\S+$")
	validateEmail := re.MatchString(u.Email)
	if !validateEmail {
		return errors.New("Invalid Email")
	}


	//Lack of SQL knowledge, there is a unique constraint
	//// since Gorm Docs doesn't seem to have a way to make emails Unique i wrote this
	//// to check the database if the email exists if it does it will give an error.
	//db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//var checkIfEmailExist user
	//db.Where("email= ?", u.Email).First(&checkIfEmailExist)
	//
	//if checkIfEmailExist.ID != u.ID {
	//	return errors.New("Email Already in Use")
	//}
	u.Password = hashedPassword
	return
}






// HashPassword is a function to hash a password will be used in auth system  using example from  https://gowebexamples.com/password-hashing/
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is a function to compare hashed passwords will be used in auth system
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
