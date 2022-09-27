package slip

import (
	"bytes"
	"io"
	"testing"
)

type rwTester struct {
	buff []byte
	r    *bytes.Reader
}

func (r *rwTester) Read(p []byte) (n int, err error) {
	if r.r == nil {
		r.r = bytes.NewReader(r.buff)
	}
	n, err = r.r.Read(p)
	if err == io.EOF {
		r.r = nil
	}
	return n, err
}

func (r *rwTester) Write(p []byte) (n int, err error) {
	r.buff = p
	r.r = nil
	return len(p), nil
}

func TestReadWriter(t *testing.T) {
	rwTester := &rwTester{}
	w := NewWriter(rwTester)
	r := NewReader(rwTester)
	payload := []byte{0xAA, 0xBB, end, esc}
	w.WritePacket(payload)
	if pkg, err := r.ReadPacket(); err == nil {
		if !compare(payload, pkg) {
			t.Fatal("What was written did not read back the same")
		}
	} else {
		t.Fatal("Received error, expected frame")
	}
}
