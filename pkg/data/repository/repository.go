package repository

import (
	"github.com/hixraid/blog/pkg/data/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(input model.UserInput) (int, error)
	Get(email string) (model.User, error)
	GetAll() ([]model.UserOutput, error)
	GetById(userId int) (model.UserOutput, error)
	UpdateById(userId int, input model.UserInput) error
	DeleteById(userId int) error
}

type PostRepository interface {
	Create(input model.PostInput) (int, error)
	GetAll() ([]model.Post, error)
	GetById(postId int) (model.Post, error)
	UpdateById(postId int, input model.PostInput) error
	DeleteById(postId int) error
}

type CommentRepository interface {
	Create(postId int, input model.CommentInput) (int, error)
	GetAll(postId int) ([]model.Comment, error)
	UpdateById(commentId int, input model.CommentInput) error
	DeleteById(commentId int) error
}

type Repository struct {
	UserRepository
	PostRepository
	CommentRepository
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:    newUserMySql(db),
		PostRepository:    newPostMySql(db),
		CommentRepository: newCommentMySql(db),
	}
}
