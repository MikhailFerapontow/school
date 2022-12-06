package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

const (
	salt      = "ghdkgn22ad4xck"
	TokenTTL  = 12 * time.Hour
	TokenSign = "qcbfn#4adfcdfe34z"
)

type tokenClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

type AuthService struct {
	repo repository.Auth
}

func newAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterAdmin(adm school.RegisterAdmin) error {
	adm.Password = s.generatePasswordHash(adm.Password)

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

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	login, err := s.repo.GetUser(login, s.generatePasswordHash(password)) //TODO generatePasswordHash(password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		login,
	})

	return token.SignedString([]byte(TokenSign))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(TokenSign), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Login, nil
}

func (s *AuthService) GetUserRole(login any) (string, error) {
	role, err := s.repo.GetUserRole(login)
	if err != nil {
		return "", errors.New("Forbid acess")
	}

	return role, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%s", hash.Sum([]byte(salt)))
}
