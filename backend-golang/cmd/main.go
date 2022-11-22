package main

import (
	"database/sql"
	"log"

	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/handler"
	"github.com/MikhailFerapontow/school/pkg/repository"
	"github.com/MikhailFerapontow/school/pkg/service"
)

var db *sql.DB

func main() {
	// fmt.Println("Hello World!")
	// db, err := sql.Open("mssql", "sqlserver://sa:Mushroom!1@localhost?database=school&connection+timeout=30")
	// if err != nil {
	// 	fmt.Println("Can't open db")
	// }

	// db.Ping()

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(school.Server)
	if err := srv.Run("8081", handler.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
