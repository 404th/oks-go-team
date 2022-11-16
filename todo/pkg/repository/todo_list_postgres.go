package repository

import (
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
