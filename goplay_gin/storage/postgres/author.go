package postgres

import (
	"fmt"

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
	query := `SELECT 
			id,
			firstname,
			secondname
		FROM author WHERE id = $1;`

	row := rp.db.QueryRow(query, id)
	if err := row.Scan(&author.ID, &author.Firstname, &author.Secondname); err != nil {
		return nil, err
	}

	return &author, nil
}

func (rp authorRepo) GetAllAuthor(offset, limit, search string) (*model.GetAllAuthor, error) {
	var (
		resp   model.GetAllAuthor
		err    error
		filter string
		params = make(map[string]interface{})
	)

	if search != "" {
		filter += " AND name ILIKE '%' || :search || '%' "
		params["search"] = search
	}

	countQuery := `SELECT count(1) FROM author WHERE true ` + filter

	err = rp.db.QueryRow(countQuery).Scan(&resp.Count)
	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	query := `SELECT
				id,
				name
			FROM author
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = limit
	params["offset"] = offset

	rows, err := rp.db.Query(query, params)
	if err != nil {
		return nil, fmt.Errorf("error while getting rows %w", err)
	}

	for rows.Next() {
		var author model.Author

		err = rows.Scan(
			&author.ID,
			&author.Firstname,
			&author.Secondname,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning author err: %w", err)
		}

		resp.Author = append(resp.Author, author)
	}

	return &resp, nil
}

func (rp authorRepo) UpdateAuthor(id string, entity model.UpdateAuthor) error {
	return nil
}

func (rp authorRepo) DeleteAuthor(id string) error {
	return nil
}
