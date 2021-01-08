package subnet

import "net"

// GetBroadcastIp 获取网段的广播地址
func GetBroadcastIp(ipNet *net.IPNet) net.IP {
	broadcastIp := make(net.IP, len(ipNet.IP))
	for i, b := range ipNet.Mask {
		broadcastIp[i] = (0xff ^ b) + ipNet.IP[i]
	}
	return broadcastIp
}
