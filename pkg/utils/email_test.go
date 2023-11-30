package utils

import "testing"

func TestIsValidEmail(t *testing.T) {
	emails := map[string]bool{
		"abcdef@gmail.com":   true,
		"ab.cd.ef@gmail.com": true,
		"abcde*@gmail.com":   false,
		"abcdef@gmail.abcde": false,
		"abcdef@gmail.a":     false,
		"abcdef@gmail.ab*":   false,
		"abcdef@gmail*.com":  false,
	}

	for k, v := range emails {
		if IsValidEmail(k) != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
