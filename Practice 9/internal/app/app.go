package app

import (
	"context"
	"pr9/config"
	"pr9/internal/delivery"
	"pr9/internal/repository"
	"pr9/internal/server"
	"pr9/internal/usecase"
	logger "pr9/pkg/logger"
	"pr9/pkg/mongoconnector"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) App {
	return App{cfg: cfg}
}

func (a *App) Start(ctx context.Context) error {

	logger := logger.InitLogger()

	mongoConn := mongoconnector.New(a.cfg.Mongo)

	repo := repository.New(mongoConn)
	uc := usecase.New(repo)

	handlers := delivery.New(a.cfg.Server, uc, logger)

	server := server.New(a.cfg, handlers, logger)

	return server.Run(ctx)
}
