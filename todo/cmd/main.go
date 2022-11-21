package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

	// getting new database
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

	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error performed while starting server: %v", err)
		}
	}()

	logrus.Println("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server is shutting down...")

	if err := srv.Close(context.Background()); err != nil {
		logrus.Printf("error occured while shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Printf("error occured while closing database: %s", err.Error())
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("config.yml")

	return viper.ReadInConfig()
}
