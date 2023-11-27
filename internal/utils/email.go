package utils

import "regexp"

const emailRegexpPattern = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"

func ValidateEmail(email string) bool {
	ok, err := regexp.MatchString(emailRegexpPattern, email)

	return ok && err == nil
}
