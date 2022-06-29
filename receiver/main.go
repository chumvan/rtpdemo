package main

import (
	"fmt"
)

// MUST be exchanged using other means
const (
	receiverAddress = "127.0.0.1"
	receiverPort    = 6420

	mtu  = 1500
	data = "Hello"
)

func main() {
	Receive(receiverAddress, receiverPort, mtu)
	fmt.Println("Finished program")
}
