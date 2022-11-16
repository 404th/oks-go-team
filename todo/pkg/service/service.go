package service

import (
	"github.com/404th/todo/model"
	"github.com/404th/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(model.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
