package slip

import "io"

const (
	end    = 0300
	esc    = 0333
	escEnd = 0334
	escEsc = 0335
)

type Reader interface {
	ReadPacket() ([]byte, error)
}

type Writer interface {
	WritePacket(data []byte) error
}

type ReadWriter interface {
	Reader
	Writer
}

type readWriter struct {
	*reader
	*writer
}

func NewReadWriter(rw io.ReadWriter) ReadWriter {
	return &readWriter{
		reader: &reader{r: rw},
		writer: &writer{o: rw},
	}
}
