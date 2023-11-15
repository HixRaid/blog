package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 16

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
