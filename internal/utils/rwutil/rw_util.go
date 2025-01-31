package rwutil

import "io"

type ReadFile func(r io.Reader) ([]byte, error)

type Reader interface {
	ReadAll(r io.Reader) ([]byte, error)
}
