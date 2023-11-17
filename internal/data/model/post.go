package model

import "time"

type Post struct {
	PostId    int       `db:"post_id" json:"post_id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostIdResponse struct {
	PostId int `json:"post_id"`
}
