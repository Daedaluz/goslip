package slip

import (
	"bytes"
	"io"
)

func DecodeFromBytes(data []byte) ([]byte, error) {
	reader := NewReader(bytes.NewReader(data))
	return reader.ReadPacket()
}

func EncodeToBytes(data []byte) []byte {
	b := &bytes.Buffer{}
	writer := NewWriter(b)
	writer.WritePacket(data)
	return b.Bytes()
}

func DecodeAllFromBytes(data []byte) ([][]byte, error) {
	reader := bytes.NewReader(data)
	return ReadAll(reader)
}

func ReadAll(r io.Reader) ([][]byte, error) {
	frames := make([][]byte, 0, 20)
	reader := NewReader(r)
	var err error
	var x []byte
	for x, err = reader.ReadPacket(); err == nil; x, err = reader.ReadPacket() {
		frames = append(frames, x)
	}
	if err != io.EOF {
		return nil, err
	}
	return frames, nil
}
