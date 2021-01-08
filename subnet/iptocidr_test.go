package subnet

import (
	"encoding/binary"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bitLength(t *testing.T) {
	assert.Equal(t, 5, bitLength(17))
	assert.Equal(t, 4, bitLength(8))

	ip := net.ParseIP("192.168.1.2")
	ipIndex := binary.BigEndian.Uint32(ip.To4())
	ipIndexMaxNum := int(ipIndex & -ipIndex)
	println(ipIndexMaxNum)
}

func Test_Ipv4AndNumToCIDR(t *testing.T) {
	ip := net.ParseIP("255.0.0.7")

	start := binary.BigEndian.Uint32(ip.To4())
	ipNets := Ipv4AndNumToCIDR(start, 10)

	assert.Equal(t, 3, len(ipNets))
	assert.Equal(t, "255.0.0.7/32", ipNets[0].String())
	assert.Equal(t, "255.0.0.8/29", ipNets[1].String())
	assert.Equal(t, "255.0.0.16/32", ipNets[2].String())

}
