package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestIsValidPassword(t *testing.T) {
	passwords := map[string]bool{
		"abcdefghijklM14":                   true,
		"abcdefghijklM 14":                  false,
		"abcdeF7":                           false,
		"abcdefghijklM":                     false,
		"abcdefghijklm14":                   false,
		"abcdefghijklm":                     false,
		"abcdefghijklmnopqrstuvwxyZ1234567": false,
	}

	for k, v := range passwords {
		if IsValidPassword(k) != v {
			t.Errorf("incorrect result for '%s'", k)
		}
	}
}

func TestHashPassword(t *testing.T) {
	password := "abcdefghijklM14"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatal("incorrect result")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Fatal("incorrect result")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "abcdefghijklM14"

	hash := "$2a$16$oj4SGrEgHVq4Uvne6rTO9O4..G0UVYu.WCjmrw5L4SZQ3yI87Yn1q"

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Fatal("incorrect result")
	}
}
