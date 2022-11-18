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
	GetAll(userId int) ([]model.TodoList, error)
	GetListById(userId, id int) (model.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input model.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db),
		TodoList:      NewTodoListRepo(db),
		TodoItem:      NewTodoItemRepo(db),
	}
}
