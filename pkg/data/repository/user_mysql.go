package repository

import (
	"github.com/hixraid/blog/pkg/data/model"
	"github.com/jmoiron/sqlx"
)

type userMySql struct {
	db *sqlx.DB
}

func newUserMySql(db *sqlx.DB) *userMySql {
	return &userMySql{db}
}

func (r *userMySql) Create(input model.UserInput) (int, error) {
	query := "INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, input.Name, input.Email, input.Password, model.UserRole)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (r *userMySql) Get(email string) (model.User, error) {
	var user model.User

	query := "SELECT * FROM users WHERE email=?"
	err := r.db.Get(&user, query, email)

	return user, err
}

func (r *userMySql) GetAll() ([]model.UserOutput, error) {
	var users []model.UserOutput

	query := "SELECT user_id, name, email, role, created_at, updated_at FROM users"
	err := r.db.Select(&users, query)

	return users, err
}

func (r *userMySql) GetById(userId int) (model.UserOutput, error) {
	var user model.UserOutput

	query := "SELECT user_id, name, email, role, created_at, updated_at FROM users WHERE user_id=?"
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *userMySql) UpdateById(userId int, input model.UserInput) error {
	query := "UPDATE users SET name=?, email=?, password=? WHERE user_id=?"
	_, err := r.db.Exec(query, input.Name, input.Email, input.Password, userId)

	return err
}

func (r *userMySql) DeleteById(userId int) error {
	query := "DELETE FROM users WHERE user_id=?"
	_, err := r.db.Exec(query, userId)

	return err
}
