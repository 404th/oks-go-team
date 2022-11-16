package repository

import (
	"github.com/404th/todo/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password_hash string) (*model.User, error)
}

type TodoList interface {
	CreateList(userId int, todoList model.TodoList) (int, error)
}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		TodoList:      NewTodoListRepo(db),
	}
}
