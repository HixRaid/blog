package utils

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLen = 8
	maxPasswordLen = 32
	cost           = 16
)

func IsValidPassword(password string) bool {
	if !IsValidText(password, minPasswordLen, maxPasswordLen) {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, s := range password {
		switch {
		case unicode.IsUpper(s):
			hasUpper = true
		case unicode.IsLower(s):
			hasLower = true
		case unicode.IsDigit(s):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
