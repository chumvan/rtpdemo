package t140writer

import (
	"io"
	"os"

	"github.com/chumvan/rtpdemo/codecs"
	"github.com/pion/rtp"
)

type (
	T140Writer struct {
		writer       io.Writer
		cachedPacket *codecs.T140Packet
	}
)

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

func (t *T140Writer) Close() error {
	if t.writer != nil {
		if closer, ok := t.writer.(io.Closer); ok {
			return closer.Close()
		}
	}
	return nil
}

func (t *T140Writer) WriteRTP(packet *rtp.Packet) error {
	if len(packet.Payload) == 0 {
		return nil
	}

	rawRTP, err := t.cachedPacket.Unmarshal(packet.Payload)
	if err != nil {
		return err
	}

	if _, err = t.writer.Write(rawRTP); err != nil {
		return err
	}
	return nil
}
