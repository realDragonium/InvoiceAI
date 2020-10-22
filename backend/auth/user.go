package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"invoiceai/database"
	"regexp"
	"time"
)

func Migrate() {
	database.DBConn.AutoMigrate(&user{})
}

//######################################################
//##            Database MODELS                       ##
//##						                          ##
//##                                                  ##
//######################################################

// User model for auth system in the future?
type user struct {
	gorm.Model
	Username string `json:"username" `
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}



//######################################################
//##            Controllers                           ##
//##						                          ##
//##                                                  ##
//######################################################


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
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Generates token and sets cookie with it for 2 hours when a user signs up then we can redirect them
	token := generateJWT(user.ID)
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(time.Minute * 120)
	c.Cookie(cookie)
	return c.Status(201).JSON(fiber.Map{
		"email":    user.Email,
		"username": user.Username,
		"token":    token,
	})

}

// UserLogin functionality to login a user and generate a JWT for them
func UserLogin(c *fiber.Ctx) error {
	//return c.JSON(fiber.Map{"success": "Logged In"})
	db := database.DBConn

	// Stores query data found from database in users.
	var users user
	user := new(user)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(err)
	}
	// checks if email provided exists in database
	if err := db.Where("email = ?", user.Email).Find(&users).Error; err !=nil{
		return c.Status(404).JSON(fiber.Map{"error":err.Error()})
	}

	// checks if password provided and password in database is the same
	checkPassword := checkPasswordHash(user.Password, users.Password)

	if users.Username != "" && checkPassword{
		// Generates token and sets cookie with it for 2 hours when a user signs up then we can redirect them
		token := generateJWT(users.ID)
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = token
		cookie.HTTPOnly = true
		cookie.Expires = time.Now().Add(time.Minute * 120)
		c.Cookie(cookie)
		return c.Status(200).JSON(fiber.Map{"username":users.Username})
	}else{
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

}

// Just a testing route to verify Auth is working
func AuthTest(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": "Logged In"})
}


//######################################################
//##            HOOKS FOR GORM                        ##
//##						                          ##
//##                                                  ##
//######################################################


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

	u.Password = hashedPassword
	return
}


//######################################################
//##            Functionality for AUTH                ##
//##						                          ##
//##                                                  ##
//######################################################


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
