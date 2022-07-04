package codecs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	p := &T140Packet{}
	if err := p.Unmarshal([]byte{}); err == nil {
		t.Fatal("Unmarshal did not error on zero length packet")
	}

	// Non-padding packet
	// valid raw packet without padding
	// 0                   1                   2                   3
	// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |V=2|P|X| CC=0  |M|   T140 PT   |       sequence number         |
	// |10 |0|0| 0000  |1| 1100100(100)|  			27023	  	       | -> 0x80, 0xe4, 0x69, 0x8f
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |                      timestamp (1000Hz)                       |
	// |				   		3653407706   					       | -> 0xd9, 0xc2, 0x93, 0xda
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |           synchronization source (SSRC) identifier            |
	// |					 	476325762							   | -> 0x1c, 0x64, 0x27, 0x82
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |                      T.140 encoded data                       |
	// +                           Hello 	            		       | -> 0x48, 0x65, 0x6c, 0x6c, 0x6f
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	rawPacket := []byte{
		0x80, 0xe4, 0x69, 0x8f,
		0xd9, 0xc2, 0x93, 0xda,
		0x1c, 0x64, 0x27, 0x82,
		0x48, 0x65, 0x6c, 0x6c, 0x6f,
	}

	parsedPacket := &T140Packet{
		Header: T140Header{
			Version:        2,
			Padding:        false,
			Extension:      false,
			CSRCCount:      0,
			Marker:         true,
			PayloadType:    100,
			SequenceNumber: 27023,
			Timestamp:      3653407706,
			SSRC:           476325762,
		},
		Payload:     rawPacket[12:],
		PaddingSize: 0x00,
	}

	if err := p.Unmarshal(rawPacket); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(p, parsedPacket) {
		fmt.Println(p.String())
		t.Errorf("TestBasic Unmarshal: got %#v, want %#v", p, parsedPacket)
	}

	if parsedPacket.MarshalSize() != len(rawPacket) {
		t.Errorf("TestBasic MarshalSize: got %#v, want %#v", parsedPacket.MarshalSize(), len(rawPacket))
	}

	byteSlice, err := p.Marshal()
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(byteSlice, rawPacket) {
		t.Errorf("TestBasic Marshal: got %#v, want %#v", byteSlice, rawPacket)
	}

	// Packet with padding
	// valid raw packet with padding
	// 0                   1                   2                   3
	// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |V=2|P|X| CC=0  |M|   T140 PT   |       sequence number         |
	// |10 |1|0| 0000  |1| 1100101(101)|  			27023	  	       | -> 0xa0, 0xe5, 0x69, 0x8f
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |                      timestamp (1000Hz)                       |
	// |				   		3653407706   					       | -> 0xd9, 0xc2, 0x93, 0xda
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |           synchronization source (SSRC) identifier            |
	// |					 	476325762							   | -> 0x1c, 0x64, 0x27, 0x82
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |                      T.140 encoded data                       |
	// +                           Hello 	            		       | -> 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, 0x03
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

	rawPacket = []byte{
		0xa0, 0xe5, 0x69, 0x8f,
		0xd9, 0xc2, 0x93, 0xda,
		0x1c, 0x64, 0x27, 0x82,
		0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x00, 0x00, 0x03,
	}
	parsedPacket = &T140Packet{
		Header: T140Header{
			Version:        2,
			Padding:        true,
			Extension:      false,
			CSRCCount:      0,
			Marker:         true,
			PayloadType:    101,
			SequenceNumber: 27023,
			Timestamp:      3653407706,
			SSRC:           476325762,
		},
		Payload:     rawPacket[12:17],
		PaddingSize: 0x03,
	}

	if err := p.Unmarshal(rawPacket); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(p, parsedPacket) {
		t.Errorf("TestBasic Unmarshal for padded packet: got %#v, want %#v", p, parsedPacket)
	}
}
