package slip

import (
	"bytes"
	"errors"
	"testing"
)

var writeTests = []test{
	{
		data:   []byte{0xAA, 0xBB},
		expect: []byte{0xAA, 0xBB, end},
		err:    nil,
	},
	{
		data:   []byte{0xAA, 0xBB, end},
		expect: []byte{0xAA, 0xBB, esc, escEnd, end},
		err:    nil,
	},
	{
		data:   []byte{0xAA, 0xBB, esc},
		expect: []byte{0xAA, 0xBB, esc, escEsc, end},
		err:    nil,
	},
	{
		data:   []byte{0xAA, 0xBB, esc, end},
		expect: []byte{0xAA, 0xBB, esc, escEsc, esc, escEnd, end},
		err:    nil,
	},
}

func TestWrite(t *testing.T) {
	for _, test := range writeTests {
		buff := &bytes.Buffer{}
		w := NewWriter(buff)
		err := w.WritePacket(test.data)
		if test.err != nil {
			if !errors.Is(err, test.err) {
				t.Fatal("Expected", test.err, "but got", err)
			}
		} else {
			if !compare(test.expect, buff.Bytes()) {
				t.Fatal("Expected", test.expect, "but got", buff.Bytes())
			}
		}
	}
}
