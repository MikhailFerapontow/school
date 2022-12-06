package repository

import (
	"fmt"

	"github.com/MikhailFerapontow/school"
	"github.com/jmoiron/sqlx"
)

type GuardianMSSql struct {
	db *sqlx.DB
}

func NewGuardianMSSql(db *sqlx.DB) *GuardianMSSql {
	return &GuardianMSSql{db: db}
}

func (r *GuardianMSSql) CreateGuardian(input school.GuardianInput) error {
	query := fmt.Sprintf("INSERT INTO Guardian (name, surname, gender, phone, email) VALUES (N'%s', N'%s', N'%s', '%s', '%s')",
		input.Name, input.Surname, input.Gender, input.Phone, input.Email)

	_, err := r.db.Query(query)

	if err != nil {
		return err
	}

	return nil
}

func (r *GuardianMSSql) GetAll() ([]school.Guardian, error) {
	var guardians []school.Guardian

	query := fmt.Sprint("SELECT * FROM Guardian")
	err := r.db.Select(&guardians, query)

	return guardians, err
}
