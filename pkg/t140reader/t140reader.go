package t140reader

import (
	"errors"
	"io"
)

const (
	t140HeaderLen = 12 // RTP packet header
	t140BlockLen  = 32 // max payload = 30 chars
	t140BufLen    = 512
)

var (
	errNilReader = errors.New("stream is nil")
)

type T140Reader struct {
	stream                io.Reader
	bytesReadSuccessfully int64
}

// NewReader returns a new T140 reader with an io.Reader input
func NewReader(in io.Reader) (*T140Reader, error) {
	if in == nil {
		return nil, errNilReader
	}
	reader := &T140Reader{
		stream:                in,
		bytesReadSuccessfully: 0,
	}
	return reader, nil
}

// ParseNewBlock reads from stream and return t140blocks payload
// and an error if there is incomplete t140block
// Return all nil values when no more frames are available
func (i *T140Reader) ParseNewBlock() (payload []byte, err error) {
	return payload, nil
}
