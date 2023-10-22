package delivery

import "github.com/gofiber/fiber/v2"

func MapRoutes(route fiber.Router, h Handler) {
	route.Post("/concurrent", h.CalculateFactorialConcurrent)
	route.Post("/linear", h.CalculateFactorialLinear)
	route.Get("/cookie", h.GetCookie)
}
