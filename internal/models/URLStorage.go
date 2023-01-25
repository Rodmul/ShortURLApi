package models

import "github.com/google/uuid"

type URLStorage struct {
	ID       uuid.UUID `db:"id"`
	ShortURL string    `db:"short"`
	LongURL  string    `db:"long"`
}
