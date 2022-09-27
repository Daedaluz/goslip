package slip

import (
	"bytes"
	"io"
)

type reader struct {
	r io.Reader
}

func (r *reader) ReadPacket() ([]byte, error) {
	buff := &bytes.Buffer{}
	b := []byte{0}
loop:
	for {
		if _, err := r.r.Read(b); err != nil {
			if err != nil {
				return nil, err
			}
		}
		switch b[0] {
		case end:
			break loop
		case esc:
			if _, err := r.r.Read(b); err != nil {
				return nil, err
			}
			switch b[0] {
			case escEnd:
				buff.WriteByte(end)
			case escEsc:
				buff.WriteByte(esc)
			}

		default:
			buff.WriteByte(b[0])
		}
	}
	return buff.Bytes(), nil
}

func NewReader(r io.Reader) Reader {
	return &reader{r: r}
}
