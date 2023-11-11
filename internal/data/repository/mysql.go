package repository

import (
	"fmt"

	"github.com/hixraid/blog/internal/config"
	"github.com/jmoiron/sqlx"
)

func NewMySql(cfg *config.DBConfig) (*sqlx.DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sqlx.Open("mysql", sourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
