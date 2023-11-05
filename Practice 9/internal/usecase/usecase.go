package usecase

import (
	"context"
	"mime/multipart"
	"pr9/internal/models"
	"pr9/internal/repository"
)

type UserUseCase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u UserUseCase) SaveFileData(ctx context.Context, file *multipart.FileHeader) (err error) {

	return u.repo.SaveFileDataToGridFS(ctx, file)
}

func (u UserUseCase) SaveFileInfo(ctx context.Context, file models.File) (err error) {

	return u.repo.SaveFileInfo(ctx, file)
}

func (u UserUseCase) GetFilesList(ctx context.Context) (files []models.File, err error) {

	return u.repo.GetFilesList(ctx)
}

func (u UserUseCase) UpdateFileByID(ctx context.Context, file models.File) (err error) {

	return u.repo.UpdateFileByID(ctx, file)
}

func (u UserUseCase) DeleteFileByID(ctx context.Context, fileID string) (err error) {

	return u.repo.DeleteFileByID(ctx, fileID)
}

func (u UserUseCase) GetFileByID(ctx context.Context, fileID string) (file models.File, err error) {

	return u.repo.GetFileByID(ctx, fileID)
}
