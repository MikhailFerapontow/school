package service

import (
	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/repository"
)

type GuardianService struct {
	repo repository.Guardian
}

func NewGuardianService(repo repository.Guardian) *GuardianService {
	return &GuardianService{repo: repo}
}

func (s *GuardianService) CreateGuardian(input school.GuardianInput) error {
	return s.repo.CreateGuardian(input)
}

func (s *GuardianService) GetAll() ([]school.Guardian, error) {
	return s.repo.GetAll()
}
