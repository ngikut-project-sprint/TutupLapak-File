package repository

import (
	"context"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
)

func (r *defaultFileRepository) AddFile(ctx context.Context, file model.File) (uint, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	uploadedFile := file

	res := r.db.WithContext(ctx).Create(&uploadedFile)
	if res.Error != nil {
		return 0, res.Error
	}

	return uploadedFile.FileID, nil
}
