package service

import "github.com/MikhailFerapontow/school/pkg/repository"

type Guardian interface {
}

type Service struct {
	Guardian
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
