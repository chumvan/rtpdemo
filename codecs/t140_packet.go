package codecs

type T140Payloader struct{}

func (p *T140Payloader) Payload(mtu uint16, payload []byte) [][]byte {
	var out [][]byte
	if payload == nil || mtu == 0 {
		return out
	}

	for len(payload) > int(mtu) {
		o := make([]byte, mtu)
		copy(o, payload[:mtu])
		payload = payload[mtu:]
		out = append(out, o)
	}

	o := make([]byte, len(payload))
	copy(o, payload)
	return append(out, o)
}

type T140Packet struct {
	Payload []byte
}

func (p *T140Packet) Unmarshal(packet []byte) ([]byte, error) {
	if packet == nil {
		return nil, errNilPacket
	} else if len(packet) < 4 { // What is the minimum payload size ?
		return nil, errShortPacket
	}

	p.Payload = packet

	return packet, nil
}
