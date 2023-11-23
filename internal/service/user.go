package service

import (
	"errors"

	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/data/repository"
	"github.com/hixraid/blog/internal/utils"
)

type UserItem struct {
	repository repository.UserRepository
}

func NewUserItem(repos repository.UserRepository) *UserItem {
	return &UserItem{repository: repos}
}

func (s *UserItem) GetAll() ([]model.UserOutput, error) {
	return s.repository.GetAll()
}

func (s *UserItem) GetById(userId int) (model.UserOutput, error) {
	return s.repository.GetById(userId)
}

func (s *UserItem) UpdateById(userId int, input model.UserInput) error {
	if !utils.ValidatePassword(input.Password) {
		return errors.New("invalid password")
	}

	return s.repository.UpdateById(userId, input)
}

func (s *UserItem) DeleteById(userId int) error {
	return s.repository.DeleteById(userId)
}
