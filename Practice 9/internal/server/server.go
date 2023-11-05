package server

import (
	"context"
	"fmt"
	"pr9/config"
	"pr9/internal/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

type Server struct {
	app      *fiber.App
	cfg      *config.Config
	handlers delivery.Handler
	logger   *zap.Logger
}

func New(cfg *config.Config, handlers delivery.Handler, logger *zap.Logger) *Server {
	return &Server{
		app:      fiber.New(),
		cfg:      cfg,
		handlers: handlers,
		logger:   logger,
	}
}

func (s *Server) Run(_ context.Context) (err error) {

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	apiGroup := s.app.Group("/api/v1")
	s.logger.Info("Map routes")
	delivery.MapRoutes(apiGroup, s.handlers)

	s.logger.Info("Starting server")
	err = s.app.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port))
	if err != nil {
		s.logger.Fatal("failed to start server, cause:[err]", zap.String("err", err.Error()))
	}

	return err
}
