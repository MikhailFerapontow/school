package main

import (
	"database/sql"
	"log"

	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/handler"
	"github.com/MikhailFerapontow/school/pkg/repository"
	"github.com/MikhailFerapontow/school/pkg/service"
	"github.com/spf13/viper"
)

var db *sql.DB

func main() {
	// fmt.Println("Hello World!")
	// db, err := sql.Open("mssql", "sqlserver://sa:Mushroom1!@localhost:1433?database=school&connection+timeout=30")
	// if err != nil {
	// 	fmt.Println("Can't open db")
	// }

	// db.Ping()

	db, err := repository.NewMSSqlDB(repository.Config{
		Host:     "localhost",
		Port:     "1433",
		UserName: "sa",
		Password: "Mushroom1!",
		DBName:   "school",
	})

	if err != nil {
		log.Fatalf("Failed to init db. Error: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("Error during config initialization. Error message: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(school.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
