package server

import (
	"context"
	"fmt"
	"pr8/config"
	"pr8/internal/deliveryy"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	app      *fiber.App
	cfg      config.Config
	handlers deliveryy.Handler
}

func New(cfg config.Config, handlers deliveryy.Handler) *Server {
	return &Server{
		app:      fiber.New(),
		cfg:      cfg,
		handlers: handlers,
	}
}

func (s *Server) Run(_ context.Context) (err error) {

	apiGroup := s.app.Group("/api/v1")
	deliveryy.MapUserRoutes(apiGroup, s.handlers)

	err = s.app.Listen(fmt.Sprint(s.cfg.Host, s.cfg.Port))
	if err != nil {
		s.cfg.Logger.Fatal("failed to start server, cause:[err]", zap.String("err", err.Error()))
	}

	return err
}
