package slip

import (
	"bytes"
	"io"
)

type writer struct {
	o io.Writer
}

func (w *writer) WritePacket(data []byte) error {
	buff := &bytes.Buffer{}
	for _, b := range data {
		switch b {
		case esc:
			buff.WriteByte(esc)
			buff.WriteByte(escEsc)
		case end:
			buff.WriteByte(esc)
			buff.WriteByte(escEnd)
		default:
			buff.WriteByte(b)
		}
	}
	buff.WriteByte(end)
	_, err := w.o.Write(buff.Bytes())
	return err
}

func NewWriter(o io.Writer) Writer {
	return &writer{o: o}
}
