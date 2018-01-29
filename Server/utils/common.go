package utils

import (
	"log"
)

// RequestContextKey - RequestContextKey
type RequestContextKey string

var (
	// RequestContextKeyUser - Context key for user
	RequestContextKeyUser = RequestContextKey("User")
)

// FailOnError - Fail on error
func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
