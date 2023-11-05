package usecase

import (
	"context"
	"mime/multipart"
	"pr9/internal/models"
)

type Usecase interface {
	SaveFileData(ctx context.Context, file *multipart.FileHeader) (err error)
	SaveFileInfo(ctx context.Context, fileI models.File) (err error)
	GetFilesList(ctx context.Context) (files []models.File, err error)
	UpdateFileByID(ctx context.Context, file models.File) (err error)
	DeleteFileByID(ctx context.Context, fileID string) (err error)
	GetFileByID(ctx context.Context, fileID string) (file models.File, err error)
}
