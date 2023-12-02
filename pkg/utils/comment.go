package utils

import (
	"errors"

	"github.com/hixraid/blog/pkg/data/model"
)

const (
	minCommentBodyLen = 1
	maxCommentBodyLen = 500
)

func ValidateCommentInput(input model.CommentInput) error {
	if !IsValidText(input.Body, minCommentBodyLen, maxCommentBodyLen) {
		return errors.New("invalid comment input body")
	}

	return nil
}
