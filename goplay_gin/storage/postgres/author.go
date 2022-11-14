package postgres

import (
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

func (rp authorRepo) CreateAuthor(id, firstname, secondname string) (string, error) {
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

func (rp authorRepo) GetAuthor(id string) (*model.Author, error) {
	var author model.Author
	query := `SELECT (
			id,
			firsname,
			secondname
		) FROM author WHERE id = $1;`

	row := rp.db.QueryRow(query, id)
	if err := row.Scan(&author); err != nil {
		return nil, err
	}

	return &author, nil
}

func (rp authorRepo) GetAllAuthor(offset, limit, search string) (*[]model.Author, error) {
	return nil, nil
}

func (rp authorRepo) UpdateAuthor(id string, entity model.UpdateAuthor) error {
	return nil
}

func (rp authorRepo) DeleteAuthor(id string) error {
	return nil
}
