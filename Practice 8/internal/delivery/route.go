package delivery

import "github.com/gofiber/fiber/v2"

func MapRoutes(route fiber.Router, h Handler) {
	route.Post("/concurrent", h.CalculateSumConcurrent)
	route.Post("/linear", h.CalculateSumLinear)
	route.Get("/cookie", h.GetCookie)
}
