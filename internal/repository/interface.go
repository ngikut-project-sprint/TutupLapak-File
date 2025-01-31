package repository

import (
	"context"

	"github.com/98prabowo/tutuplapak-file/internal/model"
)

type FileRepository interface {
	AddFile(ctx context.Context, file model.File) (uint, error)
}
