package utils

import (
	"regexp"
)

const (
	minEmailLen        = 5
	maxEmailLen        = 255
	emailRegexpPattern = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
)

func IsValidEmail(email string) bool {
	ok, err := regexp.MatchString(emailRegexpPattern, email)

	if !IsValidText(email, minEmailLen, maxEmailLen) {
		return false
	}

	return ok && err == nil
}
