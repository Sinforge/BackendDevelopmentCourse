package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CalculateSumConcurrent(ctx *fiber.Ctx) (err error)
	CalculateSumLinear(ctx *fiber.Ctx) (err error)
	GetCookie(ctx *fiber.Ctx) (err error)
}
