package local

import (
	"fmt"

	"github.com/404th/goplay_gin/model"
)

type authorRepo struct {
	db []model.Author
}

var AuthorRepo = authorRepo{}

func (r authorRepo) GetAllAuthor() ([]model.Author, error) {
	return r.db, nil
}

func (r authorRepo) CreateAuthor(id, firstname, secondname string) (model.Author, error) {
	var resp model.Author

	resp.ID = id
	resp.Firstname = firstname
	resp.Secondname = secondname

	r.db = append(r.db, resp)

	return resp, nil
}

func (r authorRepo) GetAuthorByID(id string) (*model.Author, error) {
	for _, v := range r.db {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r authorRepo) UpdateAuthor(id string, update_author model.UpdateAuthor) (string, error) {
	for i, v := range r.db {
		if v.ID == id {
			if update_author.Firstname != "" {
				r.db[i].Firstname = update_author.Firstname
			}

			if update_author.Secondname != "" {
				r.db[i].Secondname = update_author.Secondname
			}

			return id, nil
		}
	}

	return "", fmt.Errorf("author not found for updating")
}

func (r authorRepo) DeleteAuthor(id string) error {
	for i, v := range r.db {
		if v.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("not found author to be deleted")
}
