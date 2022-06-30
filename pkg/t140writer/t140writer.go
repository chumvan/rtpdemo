package t140writer

import (
	"io"
	"os"
)

type T140Writer struct {
	writer io.Writer
}

func NewWith(w io.Writer) *T140Writer {
	return &T140Writer{
		writer: w,
	}
}

func New(filename string) (*T140Writer, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return NewWith(f), nil
}
