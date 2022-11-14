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
		filter += " AND firstname ILIKE '%' || :search || '%' OR secondname ILIKE '%' || :search || '%' "
		params["search"] = search
	}

	countQuery := `SELECT count(1) FROM author WHERE true ` + filter

	q, err := rp.db.NamedQuery(countQuery, params)
	if err != nil {
		return nil, fmt.Errorf("error while scanning count %w", err)
	}

	if q.Rows.Next() {
		q.Rows.Scan(&resp.Count)
	}

	query := `SELECT
				id,
				firstname,
				secondname
			FROM author
			WHERE true` + filter

	query += " LIMIT :limit OFFSET :offset"
	params["limit"] = limit
	params["offset"] = offset

	rows, err := rp.db.NamedQuery(query, params)
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
	var (
		exists bool              = false
		params map[string]string = map[string]string{}
	)

	query1 := `SELECT EXISTS(SELECT 1 FROM author WHERE id = $1);`

	row1 := rp.db.DB.QueryRow(query1, id)
	if err := row1.Scan(&exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("not found author to be updated")
	}

	// query1 := `UPDATE author SET `
	if entity.Firstname != "" {
		params["firstname"] = entity.Firstname
	}

	if entity.Secondname != "" {
		params["secondname"] = entity.Secondname
	}

	// UPDATE weather SET temp_lo = temp_lo+1, temp_hi = temp_lo+15, prcp = DEFAULT
	// WHERE city = 'San Francisco' AND date = '2003-07-03'
	// RETURNING temp_lo, temp_hi, prcp;

	return nil
}

func (rp authorRepo) DeleteAuthor(id string) error {
	return nil
}
