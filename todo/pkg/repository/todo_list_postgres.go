package repository

import (
	"errors"
	"fmt"

	"github.com/404th/todo/model"
	"github.com/jmoiron/sqlx"
)

type TodoListRepo struct {
	db *sqlx.DB
}

func NewTodoListRepo(db *sqlx.DB) *TodoListRepo {
	return &TodoListRepo{
		db: db,
	}
}

func (tl *TodoListRepo) CreateList(userId int, todoList model.TodoList) (int, error) {
	tx, err := tl.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// 1... creating todolist
	var (
		todolistid   int
		userslistsid int
	)

	todolistquery := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`, todoListsTable)
	row1 := tx.QueryRow(todolistquery, todoList.Title, todoList.Description)
	if err = row1.Scan(&todolistid); err != nil {
		tx.Rollback()
		return 0, err
	}

	// 2... creating users lists
	userslistquery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	row2 := tx.QueryRow(userslistquery, userId, todolistid)
	if err := row2.Scan(&userslistsid); err != nil {
		tx.Rollback()
		return 0, err
	}

	return userslistsid, tx.Commit()
}

func (tl *TodoListRepo) GetAll(userId int) ([]model.TodoList, error) {
	var todolistslice []model.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1`, todoListsTable, usersListsTable)
	if err := tl.db.Select(&todolistslice, query, userId); err != nil {
		return todolistslice, err
	}

	return todolistslice, nil
}

func (tl *TodoListRepo) GetListById(userId, id int) (model.TodoList, error) {
	var todoList model.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.id = $1 AND ul.user_id = $2", todoListsTable, usersListsTable)

	if err := tl.db.Get(&todoList, query, id, userId); err != nil {
		return todoList, err
	}

	return todoList, nil
}

func (tl *TodoListRepo) Delete(userId, listId int) error {
	query := fmt.Sprintf(`DELETE FROM %s AS tl USING %s AS ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2`, todoListsTable, usersListsTable)

	resp, err := tl.db.Exec(query, userId, listId)
	if err != nil {
		return nil
	}

	num, err := resp.RowsAffected()
	if err != nil {
		return err
	}

	if num < 1 {
		return errors.New("todolist not found")
	}

	return nil
}
