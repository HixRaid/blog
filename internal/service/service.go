package service

import (
	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/data/repository"
)

type AuthService interface {
	CreateUser(input model.UserInput) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserService interface {
	GetAll() ([]model.UserOutput, error)
	GetById(userId int) (model.UserOutput, error)
	UpdateById(userId int, input model.UserInput) error
	DeleteById(userId int) error
}

type PostService interface {
	Create(input model.PostInput) (int, error)
	GetAll() ([]model.Post, error)
	GetById(postId int) (model.Post, error)
	UpdateById(postId int, input model.PostInput) error
	DeleteById(postId int) error
}

type CommentService interface {
	Create(postId int, input model.CommentInput) (int, error)
	GetAll(postId int) ([]model.Comment, error)
	UpdateById(commentId int, input model.CommentInput) error
	DeleteById(commentId int) error
}

type Service struct {
	Auth    AuthService
	User    UserService
	Post    PostService
	Comment CommentService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Auth:    NewAuthorization(repository.UserRepository),
		User:    NewUserItem(repository.UserRepository),
		Post:    NewPostItem(repository.PostRepository),
		Comment: NewCommentItem(repository.CommentRepository),
	}
}
