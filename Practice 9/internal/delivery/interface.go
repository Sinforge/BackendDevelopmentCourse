package delivery

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetFileDataByID(ctx *fiber.Ctx) error
	GetFileByID(ctx *fiber.Ctx) error
	GetFiles(ctx *fiber.Ctx) error
	UploadFile(ctx *fiber.Ctx) error
	UpdateFileByID(ctx *fiber.Ctx) error
	DeleteFileByID(ctx *fiber.Ctx) error
}
