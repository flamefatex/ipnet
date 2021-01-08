package subnet

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBroadcastIp(t *testing.T) {
	_, ipNet, _ := net.ParseCIDR("192.168.0.1/24")

	broadcastIp := GetBroadcastIp(ipNet)

	assert.Equal(t, "192.168.0.255", broadcastIp.String())
}
