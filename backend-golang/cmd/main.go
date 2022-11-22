package main

import (
	"os"

	"github.com/MikhailFerapontow/school"
	"github.com/MikhailFerapontow/school/pkg/handler"
	"github.com/MikhailFerapontow/school/pkg/repository"
	"github.com/MikhailFerapontow/school/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error during config initialization. Error message: %s", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Error during gotenv initialization. Error message: %s", err.Error())
	}

	db, err := repository.NewMSSqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.database"),
	})

	if err != nil {
		logrus.Fatalf("Failed to init db. Error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(school.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
