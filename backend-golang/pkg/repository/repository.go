package repository

import (
	"github.com/MikhailFerapontow/school"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	RegisterAdmin(school.RegisterAdmin) error
	RegisterStudent(school.RegisterStudent) error
	RegisterTeacher(school.RegisterTeacher) error
	GetUser(login, password string) (string, error)
	GetUserRole(login any) (string, error)
}

type Guardian interface {
	CreateGuardian(school.GuardianInput) error
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
