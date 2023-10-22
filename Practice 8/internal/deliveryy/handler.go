package deliveryy

import (
	"pr8/config"
	"pr8/internal/usecase"
)

type handler struct {
	uc  usecase.Usecase
	cfg config.Config
}

func New(
	cfg config.Config,
	uc usecase.Usecase,
) Handler {
	return handler{
		uc:  uc,
		cfg: cfg,
	}
}
