package repository

import (
	"fmt"

	"github.com/MikhailFerapontow/school"
	"github.com/jmoiron/sqlx"
)

type AuthMSSQL struct {
	db *sqlx.DB
}

func newAuthMSSQL(db *sqlx.DB) *AuthMSSQL {
	return &AuthMSSQL{db: db}
}

func (r *AuthMSSQL) RegisterAdmin(adm school.RegisterAdmin) error {
	query := fmt.Sprintf("EXEC Register_Admin '%s', '%s'", adm.Login, adm.Password)
	_, err := r.db.Query(query)

	if err != nil {
		return err
	}
	return nil
}

func (r *AuthMSSQL) RegisterStudent(student school.RegisterStudent) error {
	query := fmt.Sprintf("EXEC Register_Student %s, %s, %s, %s, %s, %s, %s",
		student.Login, student.Password, student.Name, student.Surname, student.Gender, student.Phone, student.Email)

	_, err := r.db.Query(query)

	if err != nil {
		return err
	}

	return nil
}

func (r *AuthMSSQL) RegisterTeacher(teacher school.RegisterTeacher) error {
	query := fmt.Sprintf("EXEC Register_Teacher %s, %s, %s, %s, %s, %s",
		teacher.Login, teacher.Password, teacher.FullName, teacher.Gender, teacher.Phone, teacher.Email)

	_, err := r.db.Query(query)
	if err != nil {
		return err
	}

	return nil
}
