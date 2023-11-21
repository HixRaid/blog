package utils

import "testing"

func TestValidatePassword(t *testing.T) {
	passwords := map[string]bool{
		"abcdefghijklM14":                   true,
		"abcdeF7":                           false,
		"abcdefghijklM":                     false,
		"abcdefghijklm14":                   false,
		"abcdefghijklm":                     false,
		"abcdefghijklmnopqrstuvwxyZ1234567": false,
	}

	for k, v := range passwords {
		if ValidatePassword(k) != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
