package codecs

// Follow standard: https://datatracker.ietf.org/doc/html/rfc4103

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

const (
	t140HeaderSize = 12 // from V to SSRC

)

//
// TODO T140 Header Implementation
//	- [x] framing
//	- [ ] newT140Header()
//	- [ ] Unmarshal()
//  - [ ] Marshal()
//	- [ ] Clone()
type T140Header struct {
	Version        uint8
	Padding        bool
	Extension      bool
	CCRCCount      uint8
	Marker         bool
	PayloadType    uint8
	SequenceNumber uint16
	Timestamp      uint32
	SSRC           uint32
}

//
// TODO T140 Packet Implementation
//
// - [ ] framing
// - [ ] String()
// - [ ] Unmarshal()
// - [ ] Marshal()
// - [ ] Clone()

type T140Packet struct {
	Header  T140Header
	Payload []byte
}
