package storage

import (
	"context"

	"github.com/404th/goplay_gin/model"
)

type StorageI interface {
	Author() AuthorI
}

type AuthorI interface {
	CreateAuthor(ctx context.Context, id, firstname, secondname string) (string, error)
	GetAuthor(ctx context.Context, id string) (*model.Author, error)
	GetAllAuthor(ctx context.Context, offset, limit, search string) (*[]model.Author, error)
	UpdateAuthor(ctx context.Context, id string, entity model.UpdateAuthor) error
	DeleteAuthor(ctx context.Context, id string) error
}
