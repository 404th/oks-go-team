package model

import "time"

type Book struct {
	ID        string    `json:"book"`
	AuthorID  string    `json:"author_id" binding:"required"`
	ISBN      string    `json:"isbn" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateBook struct {
	AuthorID string `json:"author_id" binding:"required" db:"author_id"`
	ISBN     string `json:"isbn" binding:"required" db:"isbn"`
}
