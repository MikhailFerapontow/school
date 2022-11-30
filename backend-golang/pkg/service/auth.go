package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/repository"
	"github.com/sirupsen/logrus"
)

const salt = "ghdkgn22ad4xck"

type AuthService struct {
	repo repository.Auth
}

func newAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterAdmin(adm school.RegisterAdmin) error {
	adm.Password = s.generatePasswordHash(adm.Password)
	logrus.Printf("Binded JSON, %s, %s", adm.Login, adm.Password)

	return s.repo.RegisterAdmin(adm)
}

func (s *AuthService) RegisterStudent(student school.RegisterStudent) error {
	student.Password = s.generatePasswordHash(student.Password)

	logrus.Printf("Binded Json %s, %s", student.Login, student.Password)

	return s.repo.RegisterStudent(student)
}

func (s *AuthService) RegisterTeacher(teacher school.RegisterTeacher) error {
	teacher.Password = s.generatePasswordHash(teacher.Password)

	logrus.Printf("Binded Json %s, %s", teacher.Login, teacher.Password)

	return s.repo.RegisterTeacher(teacher)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%s", hash.Sum([]byte(salt)))
}
