package controller

import (
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/service"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/validation"
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
