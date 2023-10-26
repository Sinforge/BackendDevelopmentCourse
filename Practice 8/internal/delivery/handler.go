package delivery

import (
	"encoding/json"
	"pr8/config"
	"pr8/internal/models"
	"pr8/internal/usecase"
	"pr8/utils/coder"
	"time"
	"fmt"
	"encoding/base64"


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

func (h handler) CalculateSumConcurrent(ctx *fiber.Ctx) (err error) {
	h.cfg.Logger.Info("Calculate factorial concurrent request received")

	in := models.CalculateSum{}

	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	number, err := h.uc.CalculateSumConcurrent(ctx.Context(), in.Number)
	if err != nil {
		return err
	}

	response, err := json.Marshal(number)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}
	h.cfg.Logger.Info(fmt.Sprintf("Расшифрованные данные: %s", response))
	value, err := coder.EncryptCookie(response)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}
	h.cfg.Logger.Info(fmt.Sprintf("Зашифрованные данные: %s", value))
	ctx.Cookie(&fiber.Cookie{
		Name:     h.cfg.CookieName,
		Value:    base64.StdEncoding.EncodeToString(value),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(number)
}

func (h handler) CalculateSumLinear(ctx *fiber.Ctx) (err error) {
	h.cfg.Logger.Info("Calculate factorial linear request received")

	in := models.CalculateSum{}

	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	number, err := h.uc.CalculateSumLinear(ctx.Context(), in.Number)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}

	response, err := json.Marshal(number)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}

	value, err := coder.EncryptCookie(response)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return fiber.ErrInternalServerError
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     h.cfg.CookieName,
		Value:    base64.StdEncoding.EncodeToString(value),
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(number)
}

func (h handler) GetCookie(ctx *fiber.Ctx) (err error) {
	h.cfg.Logger.Info("Get cookie request received")

	cookie := ctx.Cookies(h.cfg.CookieName)

	cipher, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	decryptedCookie, err := coder.DecryptCookie(cipher) 
	if err != nil {
		h.cfg.Logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	h.cfg.Logger.Info(fmt.Sprintf("Расшифрованные данные: %s", decryptedCookie))


	// maybe unmarshall cookie :)
	return ctx.Status(fiber.StatusOK).JSON(cookie)
}
