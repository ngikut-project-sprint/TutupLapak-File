package validation

import (
	"mime/multipart"
	"path/filepath"
	"strings"
)

type ValidateFileFunc func(fileHeader *multipart.FileHeader, maxFileSize int64) error

type FileValidator interface {
	ValidateFile(fileHeader *multipart.FileHeader, maxFileSize int64) error
}

func ValidateFile(fileHeader *multipart.FileHeader, maxFileSize int64) error {
	allowedExtensions := map[string]bool{
		".jpeg": true,
		".jpg":  true,
		".png":  true,
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return InvalidFileType
	}

	if fileHeader.Size > maxFileSize {
		return InvalidFileSize
	}

	return nil
}
