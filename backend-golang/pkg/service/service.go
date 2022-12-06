package service

import (
	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/repository"
)

type Authorization interface {
	RegisterAdmin(school.RegisterAdmin) error
	RegisterStudent(school.RegisterStudent) error
	RegisterTeacher(school.RegisterTeacher) error
	GenerateToken(username, passwrod string) (string, error)
	ParseToken(token string) (string, error)
	GetUserRole(login any) (string, error)
}

type Guardian interface {
	GetAll() ([]school.Guardian, error)
}

type Student interface {
}

type Calendar interface {
}

type Classroom interface {
}

type Grade interface {
}

type Teacher interface {
}

type Service struct {
	Authorization
	Guardian
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Auth),
		Guardian:      NewGuardianService(repos.Guardian),
	}
}
