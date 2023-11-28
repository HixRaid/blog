package utils

import (
	"errors"

	"github.com/hixraid/blog/internal/data/model"
)

func ValidateUserInput(input model.UserInput) error {
	if !IsValidName(input.Name) {
		return errors.New("invalid name")
	}

	if !IsValidEmail(input.Email) {
		return errors.New("invalid email")
	}

	if !IsValidPassword(input.Password) {
		return errors.New("invalid password")
	}

	return nil
}
