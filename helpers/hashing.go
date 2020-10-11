package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is a function to hash a password will be used in auth system  using example from  https://gowebexamples.com/password-hashing/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is a function to compare hashed passwords will be used in auth system
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
