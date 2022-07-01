package codecs

// Follow standard: https://datatracker.ietf.org/doc/html/rfc4103

// a text/t140 RTP packet without redundancy
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+ -----------------------------
// |V=2|P|X| CC=0  |M|   T140 PT   |       sequence number         |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |                      timestamp (1000Hz)                       |		Header
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           synchronization source (SSRC) identifier            |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+ ------------------------------
// |                      T.140 encoded data                       |		Payload
// +                                               +---------------+
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+				 ------------------------------

const (
	t140HeaderSize = 12 // from V to SSRC
	t140BlockSize  = 7  // ≈ T.140 encoded data size

)

//
// TODO T140 Header Implementation
//

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

type T140Packet struct {
	T140Header
	Payload []byte
}
