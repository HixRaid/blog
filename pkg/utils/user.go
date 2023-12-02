package utils

import (
	"errors"

	"github.com/hixraid/blog/pkg/data/model"
)

func ValidateUserInput(input model.UserInput) error {
	if !IsValidName(input.Name) {
		return errors.New("invalid user input name")
	}

	if !IsValidEmail(input.Email) {
		return errors.New("invalid user input email")
	}

	if !IsValidPassword(input.Password) {
		return errors.New("invalid user input password")
	}

	return nil
}
