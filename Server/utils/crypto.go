package utils

import (
	"crypto/rand"
	"fmt"
	"io"

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

// GenerateNewGUID - generates new GUID
func GenerateNewGUID() (string, error) {
	guid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, guid)

	if n != len(guid) || err != nil {
		return "AAAAAAAAAA", err
	}

	// variant bits; RFC section 4.1.1
	guid[8] = guid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); RFC section 4.1.3
	guid[6] = guid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", guid[0:4], guid[4:6], guid[6:8], guid[8:10], guid[10:]), nil
}
