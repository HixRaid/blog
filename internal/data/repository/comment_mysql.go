package repository

import (
	"github.com/hixraid/blog/internal/data/model"
	"github.com/jmoiron/sqlx"
)

type commentMySql struct {
	db *sqlx.DB
}

func newCommentMySql(db *sqlx.DB) *commentMySql {
	return &commentMySql{db}
}

func (r *commentMySql) Create(postId int, input model.CommentInput) (int, error) {
	var id int
	query := "INSERT INTO comments (body, user_id, post_id) VALUES ($1, $2, $3)"

	row := r.db.QueryRow(query, input.Body, input.UserId, input.PostId)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *commentMySql) GetAll(postId int) ([]model.Comment, error) {
	var comments []model.Comment

	query := "SELECT * FROM comments"
	err := r.db.Select(&comments, query)

	return comments, err
}

func (r *commentMySql) UpdateById(commentId int, input model.CommentInput) error {
	query := "UPDATE comments SET body=$1, user_id=$2, post_id=$3 WHERE comment_id=$4"
	_, err := r.db.Exec(query, input.Body, input.UserId, input.PostId, commentId)

	return err
}

func (r *commentMySql) DeleteById(commentId int) error {
	query := "DELETE FROM comments WHERE comment_id=$1"
	_, err := r.db.Exec(query, commentId)

	return err
}
