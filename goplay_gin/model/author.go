package model

type Author struct {
	ID         string `json:"id" db:"id"`
	Firstname  string `json:"firstname" binding:"required" db:"firstname"`
	Secondname string `json:"secondname" binding:"required" db:"secondname"`
}

type CreateAuthor struct {
	Firstname  string `json:"firstname" binding:"required" db:"firstname"`
	Secondname string `json:"secondname" binding:"required" db:"secondname"`
}

type IDTracker struct {
	ID string `json:"id" binding:"required" db:"id"`
}

type UpdateAuthor struct {
	Firstname  string `json:"firstname" db:"firstname"`
	Secondname string `json:"secondname" db:"secondname"`
}
