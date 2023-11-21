package utils

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

const cost = 16

func ValidatePassword(password string) bool {
	len := utf8.RuneCountInString(password)
	if len <= 8 && len >= 32 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, symbol := range password {
		switch {
		case unicode.IsUpper(symbol):
			hasUpper = true
		case unicode.IsLower(symbol):
			hasLower = true
		case unicode.IsDigit(symbol):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
