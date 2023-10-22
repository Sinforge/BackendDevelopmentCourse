package deliveryy

import (
	"encoding/json"
	"pr8/config"
	"pr8/internal/models"
	"pr8/internal/usecase"
	"pr8/utils/coder"
	"time"

	"github.com/gofiber/fiber/v2"
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

func (h handler) CalculateFactorialConcurrent(ctx *fiber.Ctx) (err error) {
	h.cfg.Logger.Info("Calculate factorial concurrent request received")

	in := models.CalculateFibonacci{}

	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	number, err := h.uc.CalculateFactorialConcurrent(ctx.Context(), in.Number)
	if err != nil {
		return err
	}

	response, err := json.Marshal(number)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}

	value := coder.Sha512Hash(response)

	ctx.Cookie(&fiber.Cookie{
		Name:     h.cfg.CookieName,
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(number)
}

func (h handler) CalculateFactorialLinear(ctx *fiber.Ctx) (err error) {
	h.cfg.Logger.Info("Calculate factorial linear request received")

	in := models.CalculateFibonacci{}

	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	number, err := h.uc.CalculateFactorialLinear(ctx.Context(), in.Number)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}

	response, err := json.Marshal(number)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}

	value := coder.Sha512Hash(response)

	ctx.Cookie(&fiber.Cookie{
		Name:     h.cfg.CookieName,
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(number)
}

func (h handler) GetCookie(ctx *fiber.Ctx) (err error) {
	h.cfg.Logger.Info("Get cookie request received")

	cookie := ctx.Cookies(h.cfg.CookieName)

	// maybe unmarshall cookie :)
	return ctx.Status(fiber.StatusOK).JSON(cookie)
}
