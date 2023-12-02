package utils

import "unicode/utf8"

func IsValidText(text string, minLen, maxLen int) bool {
	length := utf8.RuneCountInString(text)
	if length < minLen || length > maxLen {
		return false
	}

	return true
}
