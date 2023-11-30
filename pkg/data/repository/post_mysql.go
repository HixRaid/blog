package repository

import (
	"github.com/hixraid/blog/pkg/data/model"
	"github.com/jmoiron/sqlx"
)

type postMySql struct {
	db *sqlx.DB
}

func newPostMySql(db *sqlx.DB) *postMySql {
	return &postMySql{db}
}

func (r *postMySql) Create(input model.PostInput) (int, error) {
	query := "INSERT INTO posts (title, body) VALUES (?, ?)"
	result, err := r.db.Exec(query, input.Title, input.Body)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (r *postMySql) GetAll() ([]model.Post, error) {
	var posts []model.Post

	query := "SELECT * FROM posts"
	err := r.db.Select(&posts, query)

	return posts, err
}

func (r *postMySql) GetById(postId int) (model.Post, error) {
	var post model.Post

	query := "SELECT * FROM users WHERE post_id=?"
	err := r.db.Get(&post, query, postId)

	return post, err
}

func (r *postMySql) UpdateById(postId int, input model.PostInput) error {
	query := "UPDATE posts SET name=?, email=? WHERE post_id=?"
	_, err := r.db.Exec(query, input.Title, input.Body, postId)

	return err
}

func (r *postMySql) DeleteById(postId int) error {
	query := "DELETE FROM posts WHERE post_id=?"
	_, err := r.db.Exec(query, postId)

	return err
}
