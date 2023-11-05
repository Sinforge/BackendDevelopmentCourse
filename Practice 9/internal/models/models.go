package models

import (
	"mime/multipart"
	"time"
)

type (
	File struct {
		ID   string                `json:"id"`
		Data *multipart.FileHeader `json:"data"`
	}

	FileBD struct {
		ID         string    `bson:"_id"`
		ChunkSize  int       `bson:"chunkSize"`
		Filename   string    `bson:"filename"`
		UploadDate time.Time `bson:"uploadDate"`
	}

	FileIDRequest struct {
		FileID int `json:"file_id"`
	}

	UpdateFileRequest struct {
		FileID string `json:"file_id"`
		Name   string `json:"name"`
	}
)
