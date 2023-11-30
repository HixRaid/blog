package service

import (
	"github.com/hixraid/blog/pkg/data/model"
	"github.com/hixraid/blog/pkg/data/repository"
)

type CommentItem struct {
	repository repository.CommentRepository
}

func NewCommentItem(repos repository.CommentRepository) *CommentItem {
	return &CommentItem{repos}
}

func (s *CommentItem) Create(postId int, input model.CommentInput) (int, error) {
	return s.repository.Create(postId, input)
}

func (s *CommentItem) GetAll(postId int) ([]model.Comment, error) {
	return s.repository.GetAll(postId)
}

func (s *CommentItem) UpdateById(commentId int, input model.CommentInput) error {
	return s.repository.UpdateById(commentId, input)
}

func (s *CommentItem) DeleteById(commentId int) error {
	return s.repository.DeleteById(commentId)
}
