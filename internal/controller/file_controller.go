package controller

import (
	"github.com/98prabowo/tutuplapak-file/internal/service"
	"github.com/98prabowo/tutuplapak-file/internal/utils/validation"
)

type FileController struct {
	service      service.FileService
	maxFileSize  int64
	validateFile validation.ValidateFileFunc
}

func NewFileController(
	service service.FileService,
	maxFileSize int64,
	validateFile validation.ValidateFileFunc,
) *FileController {
	return &FileController{
		service:      service,
		maxFileSize:  maxFileSize,
		validateFile: validateFile,
	}
}
