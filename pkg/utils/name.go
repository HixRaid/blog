package utils

import "regexp"

const nameRegexpPattern = "^[\\w]{3,60}$"

func IsValidName(name string) bool {
	ok, err := regexp.MatchString(nameRegexpPattern, name)

	return ok && err == nil
}
