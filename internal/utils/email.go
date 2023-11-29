package utils

import (
	"regexp"
	"unicode/utf8"
)

const (
	maxEmailLength     = 255
	emailRegexpPattern = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
)

func IsValidEmail(email string) bool {
	ok, err := regexp.MatchString(emailRegexpPattern, email)

	length := utf8.RuneCountInString(email)
	if length > maxEmailLength {
		return false
	}

	return ok && err == nil
}
