package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mbredikhin/snippets"
	"github.com/mbredikhin/snippets/pkg/handler"
	"github.com/mbredikhin/snippets/pkg/repository"
	"github.com/mbredikhin/snippets/pkg/service"
	"github.com/robfig/cron"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize database: %s", err.Error())
	}
	rdb := repository.NewRedisDB(repository.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		logrus.Errorf("Error occured on redis init: %s", err.Error())
	}
	repos := repository.NewRepository(db, rdb)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(snippets.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	initCron(services)

	logrus.Print("Snippets app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Snippets app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on database connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initCron(services *service.Service) {
	c := cron.New()
	c.AddFunc("@midnight", func() {
		logrus.Print(time.Now().Unix())
		if err := services.Authorization.RemoveExpiredTokensFromBlacklist(time.Now().Unix()); err != nil {
			logrus.Errorf("Error occured on tokens blacklist clean up: %s", err.Error())
		}
	})

	c.Start()
}
