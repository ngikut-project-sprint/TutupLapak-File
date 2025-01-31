package service

import (
	"fmt"
)

func ErrReadBuffer(err error) error {
	return fmt.Errorf("FileService - failed to read multipart buffer: %v", err)
}

func ErrUploadImage(err error) error {
	return fmt.Errorf("FileService - failed to upload image: %v", err)
}

func ErrDecodeThumbnail(err error) error {
	return fmt.Errorf("FileService - failed to decode thumbnail: %v", err)
}

func ErrEncodeThumbnail(err error) error {
	return fmt.Errorf("FileService - failed to encode thumbnail: %v", err)
}

func ErrUploadThumbnail(err error) error {
	return fmt.Errorf("FileService - failed to upload thumbnail: %v", err)
}
