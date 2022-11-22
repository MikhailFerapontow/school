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

func (r *GuardianMSSql) GetAll() ([]school.Guardian, error) {
	var guardians []school.Guardian

	query := fmt.Sprint("SELECT * FROM Guardian")
	err := r.db.Select(&guardians, query)

	return guardians, err
}
