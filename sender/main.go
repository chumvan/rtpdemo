package main

import (
	"fmt"
	"time"
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
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				Send(senderAddress, senderPort, receiverAddress, receiverPort, []byte(data))
				fmt.Println("A packet sent at", t)
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Finished all sending")
}
