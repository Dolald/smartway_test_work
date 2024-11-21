package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dolald/smartway_test_work/configs"
	handler "github.com/Dolald/smartway_test_work/internal/handler/http"
	"github.com/Dolald/smartway_test_work/internal/repository"
	"github.com/Dolald/smartway_test_work/internal/service"
	"github.com/Dolald/smartway_test_work/server"
)

func Run() {
	cfg, err := configs.InitConfig()
	if err != nil {
		slog.Error("error initializing configs", slog.String("error", err.Error()))
	}

	db, err := repository.NewPostgresDB(&configs.DatabaseConfig{
		Host:     cfg.DataBase.Host,
		Port:     cfg.DataBase.Port,
		Username: cfg.DataBase.Username,
		DBName:   cfg.DataBase.DBName,
		SSLMode:  cfg.DataBase.SSLMode,
		Password: cfg.DataBase.Password,
	})

	if err != nil {
		slog.Error("failed to initialize db", slog.String("error", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, cfg.Handler)

	server := new(server.Server)

	go func() {
		if err := server.Run(handlers.InitRoutes(), cfg.Server); err != nil {
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
