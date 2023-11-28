package utils

import "testing"

func TestValidateName(t *testing.T) {
	names := map[string]bool{
		"abc": true,
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678": true,
		"ab": false,
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890": false,
	}

	for k, v := range names {
		if ValidateName(k) != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
