package model

import "time"

type Comment struct {
	CommentId int       `db:"comment_id" json:"comment_id"`
	Body      string    `db:"body" json:"body"`
	UserId    int       `db:"user_id" json:"user_id"`
	PostId    int       `db:"post_id" json:"post_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CommentInput struct {
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
	PostId int    `json:"post_id"`
}

type CommentIdResponse struct {
	CommentId int `json:"comment_id"`
}
