package sender

import (
	"fmt"
	"net"
)

func Send(receiverAddress string, receiverPort int, data []byte) {
	receiverAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", receiverAddress, receiverPort))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp4", nil, receiverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sender's writing")
	conn.Write(data)
	fmt.Println("Sender exit")
}
