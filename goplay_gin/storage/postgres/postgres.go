package postgres

import (
	"fmt"

	"github.com/404th/goplay_gin/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresRepo struct {
	db *sqlx.DB

	author *authorRepo
}

func NewPostgresRepo(str string) (*postgresRepo, error) {
	pg, err := sqlx.Connect("postgres", str)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to psql database: %w", err)
	}

	if err := pg.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	return &postgresRepo{
		db:     pg,
		author: &authorRepo{db: pg},
	}, nil
}

func (pg *postgresRepo) DBClose() error {
	return pg.db.Close()
}

func (pr *postgresRepo) Author() storage.AuthorI {
	return pr.author
}
