package imageutil

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/nfnt/resize"
)

type ImageCompress func(buffer *bytes.Buffer, img image.Image, maxSize int64) error

type ImageCompressor interface {
	CompressImage(buffer *bytes.Buffer, img image.Image, maxSize int64) error
}

func CompressImage(buffer *bytes.Buffer, img image.Image, maxSize int64) error {
	width := uint(img.Bounds().Dx())
	sizeReduce := uint(6)

	for {
		// Clear the buffer to avoid appending multiple times
		buffer.Reset()

		// Resize the image with reduced size (higher size reducer)
		img = resize.Resize(
			width/sizeReduce,
			0,
			img,
			resize.Lanczos3,
		)

		// Encode the image into the buffer
		err := jpeg.Encode(buffer, img, nil)
		if err != nil {
			return err
		}

		// Check if the size is under the target size
		if int64(buffer.Len()) <= maxSize {
			break
		}

		// Increate image size reducer if the image size is too large
		sizeReduce += 1
	}

	return nil
}
