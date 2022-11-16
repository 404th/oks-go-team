package service

import (
	"fmt"

	"github.com/404th/todo/model"
	"github.com/404th/todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (ts *TodoListService) CreateList(userId int, todoList model.TodoList) (int, error) {
	fmt.Println("++++++++++++++++")
	fmt.Println(userId)
	fmt.Println("++++++++++++++++")

	fmt.Println("++++++++++++++++")
	fmt.Println("++++++++++++++++")
	fmt.Println(todoList)
	fmt.Println("++++++++++++++++")

	return ts.repo.CreateList(userId, todoList)
}
