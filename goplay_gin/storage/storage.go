package storage

import (
	"github.com/404th/goplay_gin/model"
)

type StorageI interface {
	Author() AuthorI
}

type AuthorI interface {
	CreateAuthor(id, firstname, secondname string) (string, error)
	GetAuthor(id string) (*model.Author, error)
	GetAllAuthor(offset, limit, search string) (*[]model.Author, error)
	UpdateAuthor(id string, entity model.UpdateAuthor) error
	DeleteAuthor(id string) error
}
