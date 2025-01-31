package validation

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type ValidateFileFunc func(fileHeader *multipart.FileHeader, maxFileSize int64) error

type FileValidator interface {
	ValidateFile(fileHeader *multipart.FileHeader) error
}

func ValidateFile(fileHeader *multipart.FileHeader, maxFileSize int64) error {
	allowedExtensions := map[string]bool{
		".jpeg": true,
		".jpg":  true,
		".png":  true,
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return errors.New("Invalid file type. Allowed: jpeg, jpg, png")
	}

	if fileHeader.Size > maxFileSize {
		return fmt.Errorf("File size exceeds 100KiB")
	}

	return nil
}
