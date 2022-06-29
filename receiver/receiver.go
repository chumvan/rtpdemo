package main

import (
	"fmt"
	"net"
	"strings"
)

type Receiver struct {
}

func Receive(receiverAddress string, receiverPort int, mtu int) {
	serverAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", receiverAddress, receiverPort))
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp4", serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Dialed to ", conn.LocalAddr().String())

	buffer := make([]byte, mtu+100)
	for {
		pktlen, raddr, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		data := strings.TrimSpace(string(buffer[:pktlen]))
		fmt.Printf("Received: %s from %s\n", data, raddr)
	}

}
