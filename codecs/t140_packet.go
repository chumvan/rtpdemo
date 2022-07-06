package codecs

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Follow standard: https://datatracker.ietf.org/doc/html/rfc4103

// A text/t140 RTP packet with one redundant T140block RFC4103-section7.7.1
// 0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |V=2|P|X| CC=0  |M|  "RED" PT   |   sequence number of primary  |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |               timestamp of primary encoding "P"               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |           synchronization source (SSRC) identifier            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|   T140 PT   |  timestamp offset of "R"  | "R" block length  |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |0|   T140 PT   | "R" T.140 encoded redundant data              |
//    +-+-+-+-+-+-+-+-+                               +---------------+
//    +                                               |               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+     +-+-+-+-+-+
//    |                "P" T.140 encoded primary data       |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

// A text/t140 RTP packet without redundancy RFC4103-section7.7.1
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+ 	-----------------------------
// |V=2|P|X| CC=0  |M|   T140 PT   |       sequence number         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                      timestamp (1000Hz)                       |		Header = 12 bytes
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           synchronization source (SSRC) identifier            |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+ 	------------------------------
// |                      T.140 encoded data                       |		Payload = 7 bytes
// +                                               +---------------+
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+				 	------------------------------

const (
	headerLength            = 12 // from V to SSRC
	versionShift            = 6
	versionMask             = 0x3
	paddingShift            = 5
	paddingMask             = 0x1
	extensionShift          = 4
	extensionMask           = 0x1
	extensionProfileOneByte = 0xBEDE
	extensionProfileTwoByte = 0x1000
	extensionIDReserved     = 0xF
	ccMask                  = 0xF
	markerShift             = 7
	markerMask              = 0x1
	ptMask                  = 0x7F
	seqNumOffset            = 2
	seqNumLength            = 2
	timestampOffset         = 4
	timestampLength         = 4
	ssrcOffset              = 8
	ssrcLength              = 4
	csrcOffset              = 12
	csrcLength              = 4
)

//
// ------------------------------------ T140 HEADER ------------------------------------
//
type T140Header struct {
	Version        uint8
	Padding        bool
	Extension      bool
	CSRCCount      uint8
	Marker         bool
	PayloadType    uint8
	SequenceNumber uint16
	Timestamp      uint32
	SSRC           uint32
}

// Unmarshal parses a slice of bytes, store parsed result in the receiving header
// Return the number of read-byte and any error
func (h *T140Header) Unmarshal(buf []byte) (n int, err error) {
	if len(buf) < headerLength {
		return 0, fmt.Errorf("%w: %d < %d", errInsufficientLengthForAHeader, len(buf), headerLength)
	}

	h.Version = uint8(buf[0] >> versionShift & versionMask)
	h.Padding = (buf[0] >> paddingShift & paddingMask) > 0
	h.Extension = (buf[0] >> extensionShift & extensionMask) > 0 // NOTE Must be False <- no extension allowed
	if h.Extension {
		return 1, fmt.Errorf("%w", errT140NoExtensionAllowed)
	}

	h.CSRCCount = uint8(buf[0] & ccMask) // NOTE Not yet in used in RFC4103, should be 0
	if h.CSRCCount != 0 {
		return 1, fmt.Errorf("%w: got %d, expect 0", errT140CCNotZero, h.CSRCCount)
	}
	h.Marker = (buf[1] >> markerShift & markerMask) > 0
	h.PayloadType = uint8(buf[1] & ptMask)

	h.SequenceNumber = binary.BigEndian.Uint16(buf[seqNumOffset : seqNumOffset+seqNumLength])
	h.Timestamp = binary.BigEndian.Uint32(buf[timestampOffset : timestampOffset+timestampLength])
	h.SSRC = binary.BigEndian.Uint32(buf[ssrcOffset : ssrcOffset+ssrcLength])

	return headerLength, nil
}

// Marshal serializes the header into bytes
// Return a slice of bytes and any error
func (h T140Header) Marshal() (buf []byte, err error) {
	buf = make([]byte, headerLength)
	n, err := h.MarshalTo(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

// MarshalTo serializes the header and writes to a buffer
// Return number of written byte and any error
func (h T140Header) MarshalTo(buf []byte) (n int, err error) {
	if len(buf) < headerLength {
		return 0, io.ErrShortBuffer
	}

	buf[0] = h.Version << versionShift
	if h.Padding {
		buf[0] |= 1 << paddingShift
	}
	if h.Extension {
		return 1, errT140NoExtensionAllowed
	}
	buf[0] |= h.CSRCCount

	buf[1] = h.PayloadType
	if h.Marker {
		buf[1] |= 1 << markerShift
	}

	binary.BigEndian.PutUint16(buf[2:4], h.SequenceNumber)
	binary.BigEndian.PutUint32(buf[4:8], h.Timestamp)
	binary.BigEndian.PutUint32(buf[8:12], h.SSRC)

	n = headerLength
	return n, nil
}

//
// ------------------------------------ T140 PAYLOAD ------------------------------------
//

// T140Payloader payloads T140 packets
type T140Payloader struct{}

// Payload fragments an input byte slice across one/more byte slices.
// Return a slice of payload-ed byte slice
func (p *T140Payloader) Payload(mtu uint16, payload []byte) (payloads [][]byte) {
	if len(payload) == 0 {
		return payloads
	}
	// TODO If the length of payload is not zero
	return payloads
}

//
// ------------------------------------ T140 PACKET ------------------------------------
//
type T140Packet struct {
	Header      T140Header
	Payload     []byte
	PaddingSize byte
}

// String returns the representation of packet in string
func (p T140Packet) String() string {
	h := p.Header
	s := "RTP T140 PACKET:\n"

	s += fmt.Sprintf("\tVersion: %v\n", h.Version)
	s += fmt.Sprintf("\tMarker: %v\n", h.Marker)
	s += fmt.Sprintf("\tPayload Type: %d\n", h.PayloadType)
	s += fmt.Sprintf("\tSequence Number: %d\n", h.SequenceNumber)
	s += fmt.Sprintf("\tTimestamp: %d\n", h.Timestamp)
	s += fmt.Sprintf("\tSSRC: %d (%x)\n", h.SSRC, h.SSRC)
	s += fmt.Sprintf("\tPayload Length: %d\n", len(p.Payload))
	return s
}

// Unmarshal parses the input slice of bytes and stores the result in the receiving Packet
// Return packet payload and any error
func (p *T140Packet) Unmarshal(buf []byte) ([]byte, error) {
	n, err := p.Header.Unmarshal(buf)
	if err != nil {
		return nil, err
	}

	end := len(buf)
	if p.Header.Padding {
		// from RDC3550
		// The last octet of the padding contains a count of how many padding octets
		// should be ignored, including itself.
		p.PaddingSize = buf[end-1]
		end -= int(p.PaddingSize)
	}
	if end < n {
		return nil, errTooSmall
	}
	p.Payload = buf[n:end]
	return p.Payload, nil
}

// MarshalTo serializes the packet and writes to the buffer
// Return number of written byte and any error
func (p T140Packet) MarshalTo(buf []byte) (n int, err error) {
	n, err = p.Header.MarshalTo(buf)
	if err != nil {
		return 0, err
	}

	if n+len(p.Payload)+int(p.PaddingSize) > len(buf) {
		return 0, io.ErrShortBuffer
	}

	m := copy(buf[n:], p.Payload)

	if p.Header.Padding {
		buf[n+m+int(p.PaddingSize-1)] = p.PaddingSize // set the last octet to be padding length, padding value doesn't matter
	}
	return n + m + int(p.PaddingSize), nil
}

// MarshalSize returns the size of the packet once marshaled
func (p T140Packet) MarshalSize() int {
	return headerLength + len(p.Payload) + int(p.PaddingSize)
}

// Marshal serializes the packet into byte slice
// Return a byte slice or any error
func (p T140Packet) Marshal() (buf []byte, err error) {
	buf = make([]byte, p.MarshalSize())
	n, err := p.MarshalTo(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}
