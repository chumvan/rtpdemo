package t140writer

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	receiverAddress = "127.0.0.1"
	receiverPort    = 6420
	senderAddress   = "127.0.0.1"
	senderPort      = 6421
	mtu             = 1500
)

var localAddrString = fmt.Sprintf("%s:%d", senderAddress, senderPort)
var remoteAddrString = fmt.Sprintf("%s:%d", receiverAddress, receiverPort)

var udp = &net.UDPConn{}

func init() {
	senderAddr, err := net.ResolveUDPAddr("udp4", localAddrString)
	if err != nil {
		panic(err)
	}
	receiverAddr, err := net.ResolveUDPAddr("udp4", remoteAddrString)
	if err != nil {
		panic(err)
	}
	udp, err = net.DialUDP("udp4", senderAddr, receiverAddr)
	if err != nil {
		panic(err)
	}
}

func TestServerInit(t *testing.T) {

	if udp.LocalAddr().Network() != "udp" ||
		udp.LocalAddr().String() != localAddrString ||
		udp.RemoteAddr().String() != remoteAddrString {
		t.Errorf("Wrong connection setting: \nExpect: ,%#v,\nGot: ,%#v,", "127.0.0.1:6421", udp.LocalAddr().String())
		udp.Close()
	}
}

func TestNewWith(t *testing.T) {
	t140Writer := NewWith(udp)
	assert.NotNil(t, t140Writer.Close)
	t140Writer.Close()
}
