package model

import (
	"errors"
	"invoiceai/helpers"
	"regexp"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model for auth system in the future?
type User struct {
	gorm.Model
	Username string `json:"username" `
	Email    string `json:"email"`

	Password string `json:"password"`
}

// BeforeCreate is a function to add password hashing and User validation see https://gorm.io/docs/hooks.html
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	// Hasing of Password before saving to database
	hashedPassword, err := helpers.HashPassword(u.Password)
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

	// since Gorm Docs doesn't seem to have a way to make emails Unique i wrote this
	// to check the database if the email exists if it does it will give an error.
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var checkIfEmailExist User
	db.Where("email= ?", u.Email).First(&checkIfEmailExist)

	if checkIfEmailExist.ID != u.ID {
		return errors.New("Email Already in Use")
	}
	u.Password = hashedPassword
	return
}
