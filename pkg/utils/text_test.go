package utils

import "testing"

func TestIsValidText(t *testing.T) {
	texts := map[string]bool{
		"abcdefghijklmn":            true,
		"abcdefghijklmnopqrstuvwx":  true,
		"abcd":                      true,
		"abc":                       false,
		"abcdefghijklmnopqrstuvwxy": false,
	}

	const (
		minLen = 4
		maxLen = 24
	)

	for k, v := range texts {
		if IsValidText(k, minLen, maxLen) != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}
