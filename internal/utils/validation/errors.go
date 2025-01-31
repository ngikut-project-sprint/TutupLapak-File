package validation

import "errors"

var (
	InvalidFileType = errors.New("Invalid file type. Allowed: jpeg, jpg, png")
	InvalidFileSize = errors.New("File size exceeds 100KiB")
)
