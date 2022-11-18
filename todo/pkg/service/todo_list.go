package service

import (
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
	return ts.repo.CreateList(userId, todoList)
}

func (ts *TodoListService) GetAll(userId int) ([]model.TodoList, error) {
	return ts.repo.GetAll(userId)
}

func (ts *TodoListService) GetListById(userId, id int) (model.TodoList, error) {
	return ts.repo.GetListById(userId, id)
}

func (ts *TodoListService) Delete(userId, listId int) error {
	return ts.repo.Delete(userId, listId)
}

func (ts *TodoListService) Update(userId, listId int, input model.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return ts.repo.Update(userId, listId, input)
}
