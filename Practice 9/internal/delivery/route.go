package delivery

import "github.com/gofiber/fiber/v2"

func MapRoutes(route fiber.Router, h Handler) {
	route.Get("/files/:id/info", h.GetFileDataByID)
	route.Get("/files/:id", h.GetFileByID)
	route.Get("/files", h.GetFiles)

	route.Post("/files", h.UploadFile)
	route.Put("/files/:id", h.UpdateFileByID)
	route.Delete("/files/:id", h.DeleteFileByID)
}
