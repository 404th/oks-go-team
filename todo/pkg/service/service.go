package service

import (
	"github.com/404th/todo/model"
	"github.com/404th/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	CreateList(userId int, todoList model.TodoList) (int, error)
}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
