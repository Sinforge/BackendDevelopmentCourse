package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CalculateFactorialConcurrent(ctx *fiber.Ctx) (err error)
	CalculateFactorialLinear(ctx *fiber.Ctx) (err error)
	GetCookie(ctx *fiber.Ctx) (err error)
}
