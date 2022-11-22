package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
}

func NewMSSqlDB(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30",
		cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}