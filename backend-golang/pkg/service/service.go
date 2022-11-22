package service

import (
	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/repository"
)

type Guardian interface {
	GetAll() ([]school.Guardian, error)
}

type Service struct {
	Guardian
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Guardian: NewGuardianService(repos.Guardian),
	}
}
