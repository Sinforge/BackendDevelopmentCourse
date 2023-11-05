package delivery

import (
	"io"
	"pr9/internal/models"
	"pr9/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type handler struct {
	uc     usecase.Usecase
	cfg    Config
	logger *zap.Logger
}

func New(
	cfg Config,
	uc usecase.Usecase,
	logger *zap.Logger,
) Handler {
	return handler{
		uc:     uc,
		cfg:    cfg,
		logger: logger,
	}
}

func (h handler) GetFileDataByID(ctx *fiber.Ctx) error {
	h.logger.Info("Get file data by id request received")

	fileID := ctx.Params("id")

	file, err := h.uc.GetFileByID(ctx.Context(), fileID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(file)
}

func (h handler) GetFileByID(c *fiber.Ctx) error {
	h.logger.Info("Get file by id request received")

	fileID := c.Params("id")

	file, err := h.uc.GetFileByID(c.Context(), fileID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Set("Content-Disposition", "attachment; filename=alohadance")
	c.Set("Content-Type", "application/octet-stream")

	fileData, err := file.Data.Open()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	_, err = io.Copy(c.Response().BodyWriter(), fileData)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)

}

func (h handler) GetFiles(ctx *fiber.Ctx) error {
	h.logger.Info("Get files request received")

	files, err := h.uc.GetFilesList(ctx.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.Status(fiber.StatusOK).JSON(files)
}

func (h handler) UploadFile(ctx *fiber.Ctx) error {
	h.logger.Info("upload file request received")

	in := models.FileIDRequest{}

	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return h.uc.SaveFileData(ctx.Context(), file)
}

func (h handler) UpdateFileByID(ctx *fiber.Ctx) error {
	h.logger.Info("update file by id request received")

	in := models.UpdateFileRequest{}

	if err := ctx.BodyParser(&in); err != nil {
		return fiber.ErrBadRequest
	}

	fileID := ctx.Params("id")

	file, err := ctx.FormFile("file")
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return h.uc.UpdateFileByID(ctx.Context(), models.File{
		ID:   fileID,
		Data: file,
	})
}

func (h handler) DeleteFileByID(ctx *fiber.Ctx) error {
	h.logger.Info("delete file by id request received")

	fileID := ctx.Params("id")

	return h.uc.DeleteFileByID(ctx.Context(), fileID)
}
