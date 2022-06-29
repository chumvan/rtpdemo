package main

import (
	"fmt"
	"net"
)

func Send(senderAddress string, senderPort int, receiverAddress string, receiverPort int, data []byte) {
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
	fmt.Println("Sender's writing")
	conn.Write(data)
	fmt.Println("Sender exit")
}
