package service

import (
	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/data/repository"
)

type PostItem struct {
	repository repository.PostRepository
}

func NewPostItem(repos repository.PostRepository) *PostItem {
	return &PostItem{repos}
}

func (s *PostItem) Create(input model.PostInput) (int, error) {
	return s.repository.Create(input)
}

func (s *PostItem) GetAll() ([]model.Post, error) {
	return s.repository.GetAll()
}

func (s *PostItem) GetById(postId int) (model.Post, error) {
	return s.repository.GetById(postId)
}

func (s *PostItem) UpdateById(postId int, input model.PostInput) error {
	return s.repository.UpdateById(postId, input)
}

func (s *PostItem) DeleteById(postId int) error {
	return s.repository.DeleteById(postId)
}
