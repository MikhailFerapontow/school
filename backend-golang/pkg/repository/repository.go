package repository

import (
	"github.com/MikhailFerapontow/school"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	RegisterAdmin(school.RegisterAdmin) error
	RegisterStudent(school.RegisterStudent) error
	RegisterTeacher(school.RegisterTeacher) error
}

type Guardian interface {
	GetAll() ([]school.Guardian, error)
}

type Repository struct {
	Auth
	Guardian
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth:     newAuthMSSQL(db),
		Guardian: NewGuardianMSSql(db),
	}
}
