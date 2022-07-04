package codecs

import "testing"

func TestBasic(t *testing.T) {
	p := &T140Packet{}
	if err := p.Unmarshal([]byte{}); err == nil {
		t.Fatal("Unmarshal did not error on zero length packet")
	}
}
