package postgres

import (
	"context"

	"github.com/404th/goplay_gin/model"
	"github.com/jmoiron/sqlx"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) *authorRepo {
	return &authorRepo{
		db: db,
	}
}

func (rp authorRepo) CreateAuthor(ctx context.Context, id, firstname, secondname string) (string, error) {
	var (
		created_id string
	)
	query := `INSERT INTO author (
		id, firstname, secondname
	) VALUES (
		$1, $2, $3
	) RETURNING id;`

	row := rp.db.QueryRow(query, id, firstname, secondname)
	if err := row.Scan(&created_id); err != nil {
		return "", err
	}

	return created_id, nil
}

func (rp authorRepo) GetAuthor(ctx context.Context, id string) (*model.Author, error) {
	return nil, nil
}

func (rp authorRepo) GetAllAuthor(ctx context.Context, offset, limit, search string) (*[]model.Author, error) {
	return nil, nil
}

// sdsds

func (rp authorRepo) UpdateAuthor(ctx context.Context, id string, entity model.UpdateAuthor) error {
	return nil
}

func (rp authorRepo) DeleteAuthor(ctx context.Context, id string) error {
	return nil
}
