package codecs

import (
	"errors"
	"testing"
)

func TestT140Packet_Unmarshal(t *testing.T) {
	pck := T140Packet{}
	// Nil packet
	raw, err := pck.Unmarshal(nil)
	if raw != nil {
		t.Fatal("Result should be nil in case of error")
	}
	if !errors.Is(err, errNilPacket) {
		t.Fatal("Error should be:", errNilPacket)
	}

	// Nil payload
	raw, err = pck.Unmarshal([]byte{})
	if raw != nil {
		t.Fatal("Result should be nil in case of error")
	}
	if !errors.Is(err, errShortPacket) {
		t.Fatal("Error should be:", errShortPacket)
	}

	// Payload smaller than header size
	raw, err = pck.Unmarshal([]byte{0x00, 0x11, 0x22})
	if raw != nil {
		t.Fatal("Result should be nil in case of error")
	}
	if !errors.Is(err, errShortPacket) {
		t.Fatal("Error should be:", errShortPacket)
	}

	// Normal payload
	normal_payload := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x90}
	raw, err = pck.Unmarshal(normal_payload)
	if raw == nil {
		t.Fatal("Result shouldn't be nil in case of success")
	}
	if err != nil {
		t.Fatal("Error should be nil in case of success")
	}

}
