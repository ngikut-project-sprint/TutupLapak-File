package service

import (
	"context"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/rwutil"
)

type FileService interface {
	UploadFile(ctx context.Context, file rwutil.FileOpener, fileName string, completion chan model.Completion)
	GenerateThumbnail(ctx context.Context, file rwutil.FileOpener, fileName string, completion chan model.Completion)
	AddFile(ctx context.Context, fileURI string, fileThumbnailURI string) (uint, error)
}
