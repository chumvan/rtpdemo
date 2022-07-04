package main

import (
	"fmt"
	"net"
)

// Send sends a byte slice to the receiver address
func Send(senderAddress string, senderPort int, receiverAddress string, receiverPort int, data []byte) error {
	receiverAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", receiverAddress, receiverPort))
	if err != nil {
		panic(err)
	}
	senderAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", senderAddress, senderPort))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp4", senderAddr, receiverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write(data)
	return nil
}
