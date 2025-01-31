package rwutil

import (
	"io"
	"mime/multipart"
)

type ReadFile func(r io.Reader) ([]byte, error)

type Reader interface {
	ReadAll(r io.Reader) ([]byte, error)
}

type FileOpener interface {
	Open() (multipart.File, error)
}

type FileData interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}
