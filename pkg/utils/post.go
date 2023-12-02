package utils

import (
	"errors"

	"github.com/hixraid/blog/pkg/data/model"
)

const (
	minPostTitleLen = 5
	maxPostTitleLen = 200
	minPostBodyLen  = 100
	maxPostBodyLen  = 5000
)

func ValidatePostInput(input model.PostInput) error {
	if !IsValidText(input.Title, minPostTitleLen, maxPostTitleLen) {
		return errors.New("invalid post input title")
	}

	if !IsValidText(input.Body, minPostBodyLen, maxPostBodyLen) {
		return errors.New("invalid post input body")
	}

	return nil
}
