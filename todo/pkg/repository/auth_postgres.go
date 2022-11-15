package repository

import (
	"github.com/404th/todo/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (ar AuthRepo) CreateUser(model.User) (int, error) {
	return 0, nil
}
