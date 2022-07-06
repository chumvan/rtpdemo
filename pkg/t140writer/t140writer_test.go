package t140writer

import (
	"fmt"
	"net"
	"testing"
)

const (
	receiverAddress = "127.0.0.1"
	receiverPort    = 6420
	senderAddress   = "127.0.0.1"
	senderPort      = 6421
	mtu             = 1500
)

var udp = &net.UDPConn{}

func init() {
	receiverAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", receiverAddress, receiverPort))
	if err != nil {
		panic(err)
	}
	senderAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", senderAddress, senderPort))
	if err != nil {
		panic(err)
	}
	udp, err = net.DialUDP("udp4", senderAddr, receiverAddr)
	if err != nil {
		panic(err)
	}
}

func TestServer(t *testing.T) {
	if udp.LocalAddr().Network() != "udp" || udp.LocalAddr().String() != "127.0.0.1:6421" {
		t.Errorf("Wrong connection setting: \nExpect: ,%#v,\nGot: ,%#v,", "127.0.0.1:6421", udp.LocalAddr().String())
	}
}
