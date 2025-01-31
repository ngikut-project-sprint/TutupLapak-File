package service

import (
	"context"
	"mime/multipart"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
)

type FileService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader, fileName string, completion chan model.Completion)
	GenerateThumbnail(ctx context.Context, file *multipart.FileHeader, fileName string, completion chan model.Completion)
	AddFile(ctx context.Context, fileURI string, fileThumbnailURI string) (uint, error)
}
