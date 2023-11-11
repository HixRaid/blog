package repository

import (
	"github.com/hixraid/blog/internal/data/model"
	"github.com/jmoiron/sqlx"
)

type postMySql struct {
	db *sqlx.DB
}

func newPostMySql(db *sqlx.DB) *postMySql {
	return &postMySql{db}
}

func (r *postMySql) Create(input model.PostInput) (int, error) {
	var id int
	query := "INSERT INTO posts (title, body) VALUES ($1, $2)"

	row := r.db.QueryRow(query, input.Title, input.Body)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *postMySql) GetAll() ([]model.Post, error) {
	var posts []model.Post

	query := "SELECT * FROM posts"
	err := r.db.Select(&posts, query)

	return posts, err
}

func (r *postMySql) GetById(postId int) (model.Post, error) {
	var post model.Post

	query := "SELECT * FROM users WHERE post_id=$1"
	err := r.db.Get(&post, query, postId)

	return post, err
}

func (r *postMySql) UpdateById(postId int, input model.PostInput) error {
	query := "UPDATE posts SET name=$1, email=$2 WHERE post_id=$3"
	_, err := r.db.Exec(query, input.Title, input.Body, postId)

	return err
}

func (r *postMySql) DeleteById(postId int) error {
	query := "DELETE FROM posts WHERE post_id=$1"
	_, err := r.db.Exec(query, postId)

	return err
}
