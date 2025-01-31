package service

import (
	"context"

	"github.com/98prabowo/tutuplapak-file/internal/model"
	"github.com/98prabowo/tutuplapak-file/internal/utils/errorutil"
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
