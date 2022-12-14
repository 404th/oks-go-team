package repository

import (
	"fmt"

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

func (ar AuthRepo) CreateUser(user model.User) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id;`, usersTable)
	row := ar.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (ar AuthRepo) GetUser(username, password_hash string) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	row := ar.db.QueryRow(query, username, password_hash)
	if err := row.Scan(&user.Id); err != nil {
		return nil, err
	}

	return &user, nil
}
