package service

import (
	"context"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/errorutil"
)

func (s *s3FileService) AddFile(ctx context.Context, fileURI string, fileThumbnailURI string) (uint, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	fileID, err := s.repo.AddFile(ctx, model.File{
		FileURI:          fileURI,
		FileThumbnailURI: fileThumbnailURI,
	})
	if err != nil {
		return 0, errorutil.ErrWithContext(err)
	}

	return fileID, nil
}
