package repository

import (
	"context"
	"mime/multipart"
	"pr9/internal/models"
)

type Repository interface {
	SaveFileDataToGridFS(ctx context.Context, file *multipart.FileHeader) (err error)
	SaveFileInfo(ctx context.Context, file models.File) (err error)
	GetFilesList(ctx context.Context) (files []models.File, err error)
	UpdateFileByID(ctx context.Context, file models.File) (err error)
	DeleteFileByID(ctx context.Context, fileID string) (err error)
	GetFileByID(ctx context.Context, fileID string) (file models.File, err error)
}
