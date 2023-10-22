package app

import (
	"context"
	"pr8/config"
	"pr8/internal/delivery"
	"pr8/internal/server"
	"pr8/internal/usecase"
)

type App struct {
	cfg config.Config
}

func NewApp(cfg config.Config) App {
	return App{cfg: cfg}
}

func (a *App) Start(ctx context.Context) error {
	// map handlers and the rest of the shit

	// repo
	uc := usecase.New()

	handlers := delivery.New(a.cfg, uc)

	server := server.New(a.cfg, handlers)

	return server.Run(ctx)
}
