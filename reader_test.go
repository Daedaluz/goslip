package slip

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

type test struct {
	data   []byte
	expect []byte
	err    error
}

var readTests = []test{
	{
		data:   []byte{},
		expect: []byte{},
		err:    io.EOF,
	},
	{
		data:   []byte{0xAA, 0xBB},
		expect: []byte{},
		err:    io.EOF,
	},
	{
		data:   []byte{0xAA, 0xBB, end},
		expect: []byte{0xAA, 0xBB},
		err:    nil,
	},
	{
		data:   []byte{0xAA, 0xBB, esc, escEnd, end},
		expect: []byte{0xAA, 0xBB, end},
		err:    nil,
	},
	{
		data:   []byte{0xAA, 0xBB, esc, escEsc, end},
		expect: []byte{0xAA, 0xBB, esc},
		err:    nil,
	},
	{
		data:   []byte{0xAA, 0xBB, esc, escEsc, esc, escEnd, end},
		expect: []byte{0xAA, 0xBB, esc, end},
		err:    nil,
	},
}

func compare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}

func TestRead(t *testing.T) {
	for _, test := range readTests {
		r := NewReader(bytes.NewReader(test.data))
		pkg, err := r.ReadPacket()
		if test.err != nil {
			if !errors.Is(err, test.err) {
				t.Fatal("Expected", test.err, "but got", err)
			}
		} else {
			if !compare(test.expect, pkg) {
				t.Fatal("Expected", test.expect, "but got", pkg)
			}
		}
	}
}
