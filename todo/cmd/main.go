package main

import (
	"os"

	"github.com/404th/todo/pkg/handler"
	"github.com/404th/todo/pkg/repository"
	"github.com/404th/todo/pkg/service"
	"github.com/404th/todo/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// set json format to logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// loading configs/config.yml
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing configs: %s", err.Error())
	}

	// loading .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while initializing global env: %s", err.Error())
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("cannot connect to database %s", err.Error())
	}
	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error performed while starting server: %v", err)
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("config.yml")

	return viper.ReadInConfig()
}
