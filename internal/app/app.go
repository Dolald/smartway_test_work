package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	handler "github.com/Dolald/smartway_test_work/internal/handler/http"
	"github.com/Dolald/smartway_test_work/internal/repository"
	"github.com/Dolald/smartway_test_work/internal/service"
	"github.com/Dolald/smartway_test_work/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Run() {

	if err := initConfig(); err != nil {
		slog.Error("error initializing configs", slog.String("error", err.Error()))
	}

	if err := godotenv.Load(); err != nil {
		slog.Error("error loading env variables: %s", slog.String("error", err.Error()))
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		slog.Error("failed to initialize db", slog.String("error", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(server.Server)

	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			slog.Error("error occured while runnung http server", slog.String("error", err.Error()))
		}
	}()

	slog.Info("web-server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	slog.Info("web-server shuttong down")

	if err := server.Shutdown(context.Background()); err != nil {
		slog.Error("error occured on server shutting down", slog.String("error", err.Error()))
	}

	if err := db.Close(); err != nil {
		slog.Error("error occured on db connection close", slog.String("error", err.Error()))
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
