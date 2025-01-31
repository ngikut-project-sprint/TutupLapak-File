package repository

import (
	"context"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
)

type FileRepository interface {
	AddFile(ctx context.Context, file model.File) (uint, error)
}
