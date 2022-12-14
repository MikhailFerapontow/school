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
	query := fmt.Sprintf("EXEC Register_Student '%s', '%s', N'%s', N'%s', N'%s', '%s', '%s'",
		student.Login, student.Password, student.Name, student.Surname, student.Gender, student.Phone, student.Email)

	_, err := r.db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthMSSQL) RegisterTeacher(teacher school.RegisterTeacher) error {
	query := fmt.Sprintf("EXEC Register_Teacher '%s', '%s', N'%s', N'%s', '%s', '%s'",
		teacher.Login, teacher.Password, teacher.FullName, teacher.Gender, teacher.Phone, teacher.Email)

	_, err := r.db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthMSSQL) GetUser(login, password string) (string, error) {
	var user string
	query := fmt.Sprintf("SELECT login FROM [User] WHERE login = '%s' AND password = '%s'", login, password)

	err := r.db.Get(&user, query)

	return user, err
}

func (r *AuthMSSQL) GetUserRole(login any) (string, error) {
	var role string
	query := fmt.Sprintf("SELECT role_name FROM User_role JOIN [User] ON [User].role_id = User_role.role_id WHERE [User].login = '%s'", login)

	err := r.db.Get(&role, query)

	return role, err
}
