package controller

import (
	"time"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/service"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/validation"
)

type FileController struct {
	service      service.FileService
	maxFileSize  int64
	validateFile validation.ValidateFileFunc
	timeout      time.Duration
}

func NewFileController(
	service service.FileService,
	maxFileSize int64,
	validateFile validation.ValidateFileFunc,
	timeout time.Duration,
) *FileController {
	return &FileController{
		service:      service,
		maxFileSize:  maxFileSize,
		validateFile: validateFile,
		timeout:      timeout,
	}
}
