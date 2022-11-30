package repository

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const (
	UserTable            = "User"
	UserRoleTable        = "User_role"
	StudentTable         = "Student"
	GuardianTable        = "Guardian"
	TeacherTable         = "Teacher"
	CalendarTable        = "Calendar"
	SubjectTable         = "Subject"
	ClassroomTable       = "Classroom"
	GradeTable           = "Grade"
	StudentGuardianTable = "student_guardian"
	StudentCalendarTable = "student_calendar"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
}

func NewMSSqlDB(cfg Config) (*sqlx.DB, error) {

	db, err := sqlx.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30",
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
