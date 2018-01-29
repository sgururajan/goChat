package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - hashes given password string
func HashPassword(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 15)
	return string(bytes), err
}

// CheckPasswordHash - Checks password has
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
