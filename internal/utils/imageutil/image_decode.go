package imageutil

import (
	"image"
	"io"
)

type DecodeImage func(r io.Reader) (image.Image, error)

type ImageDecoder interface {
	Decode(r io.Reader) (image.Image, error)
}
