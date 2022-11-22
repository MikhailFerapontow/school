package repository

import (
	"github.com/MikhailFerapontow/school"
	"github.com/jmoiron/sqlx"
)

type Guardian interface {
	GetAll() ([]school.Guardian, error)
}

type Repository struct {
	Guardian
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Guardian: NewGuardianMSSql(db),
	}
}
