package model

import (
	"gorm.io/gorm"
)

// User model for auth system in the future?
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
