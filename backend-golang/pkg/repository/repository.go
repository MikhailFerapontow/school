package repository

import (
	"database/sql"
)

type Guardian interface {
}

type Repository struct {
	Guardian
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
