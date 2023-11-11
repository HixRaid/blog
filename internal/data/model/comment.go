package model

import "time"

type Comment struct {
	CommentId int       `db:"comment_id"`
	Body      string    `db:"body"`
	UserId    int       `db:"user_id"`
	PostId    int       `db:"post_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CommentInput struct {
	Body   string `json:"body"`
	UserId int    `json:"user_id"`
	PostId int    `json:"post_id"`
}
