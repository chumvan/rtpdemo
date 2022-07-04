package main

import (
	"fmt"
	"time"

	"github.com/chumvan/rtpdemo/codecs"
)

// MUST be changed using other means
const (
	receiverAddress = "127.0.0.1"
	receiverPort    = 6420
	senderAddress   = "127.0.0.1"
	senderPort      = 6421
	mtu             = 1500
	data            = "Hello"
)

func main() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)
	rawPacket := []byte{
		0x80, 0xe4, 0x69, 0x8f,
		0xd9, 0xc2, 0x93, 0xda,
		0x1c, 0x64, 0x27, 0x82,
		0x48, 0x65, 0x6c, 0x6c, 0x6f,
	}
	parsedPacket := &codecs.T140Packet{
		Header: codecs.T140Header{
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
	rawRTP, err := parsedPacket.Marshal()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				Send(senderAddress, senderPort, receiverAddress, receiverPort, rawRTP)
				fmt.Println("A packet sent at", t)
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Finished all sending")
}
