package subnet

import (
	"encoding/binary"
	"net"
)

//Given a start IP address ip and a number of ips we need to cover n, return a representation of the range as a list (of smallest possible length) of CIDR blocks.
//
//A CIDR block is a string consisting of an IP, followed by a slash, and then the prefix length. For example: "123.45.67.89/20". That prefix length "20" represents the number of common prefix bits in the specified range.
//
//Example 1:
//Input: ip = "255.0.0.7", n = 10
//Output: ["255.0.0.7/32","255.0.0.8/29","255.0.0.16/32"]
//Explanation:
//The initial ip address, when converted to binary, looks like this (spaces added for clarity):
//255.0.0.7 -> 11111111 00000000 00000000 00000111
//The address "255.0.0.7/32" specifies all addresses with a common prefix of 32 bits to the given address,
//ie. just this one address.
//
//The address "255.0.0.8/29" specifies all addresses with a common prefix of 29 bits to the given address:
//255.0.0.8 -> 11111111 00000000 00000000 00001000
//Addresses with common prefix of 29 bits are:
//11111111 00000000 00000000 00001000
//11111111 00000000 00000000 00001001
//11111111 00000000 00000000 00001010
//11111111 00000000 00000000 00001011
//11111111 00000000 00000000 00001100
//11111111 00000000 00000000 00001101
//11111111 00000000 00000000 00001110
//11111111 00000000 00000000 00001111
//
//The address "255.0.0.16/32" specifies all addresses with a common prefix of 32 bits to the given address,
//ie. just 11111111 00000000 00000000 00010000.
//
//In total, the answer specifies the range of 10 ips starting with the address 255.0.0.7 .
//
//There were other representations, such as:
//["255.0.0.7/32","255.0.0.8/30", "255.0.0.12/30", "255.0.0.16/32"],
//but our answer was the shortest possible.
//
//Also note that a representation beginning with say, "255.0.0.7/30" would be incorrect,
//because it includes addresses like 255.0.0.4 = 11111111 00000000 00000000 00000100
//that are outside the specified range.
//Note:
//ip will be a valid IPv4 address.
//Every implied address ip + x (for x < n) will be a valid IPv4 address.
//n will be an integer in the range [1, 1000].

// Ipv4AndNumToCIDR ip范围到子网掩码格式的转换，
// ip从start开始，取n个ip,得到的网段
func Ipv4AndNumToCIDR(start uint32, n int) (ipNets []*net.IPNet) {
	ipNets = make([]*net.IPNet, 0)

	var mask int
	ipIndex := start

	for n > 0 {

		ipIndexMaxNum := int(ipIndex & -ipIndex)
		// ipIndexMaxNum，此ipIndex下，所能最多的ip数量
		// 	ip := net.ParseIP("192.168.1.2")
		//	ipIndex := binary.BigEndian.Uint32(ip.To4())
		//	ipIndexMaxNum := int(ipIndex & -ipIndex)
		// 如ip=192.168.1.0，ipIndexMaxNum=256
		// 如ip=192.168.1.1，ipIndexMaxNum=1
		// 如ip=192.168.1.2，ipIndexMaxNum=2

		// 取掩码大的
		ipIndexMaxNumMask := 33 - bitLength(ipIndexMaxNum)
		nMaxMask := 33 - bitLength(n)
		if ipIndexMaxNumMask > nMaxMask {
			mask = ipIndexMaxNumMask
		} else {
			mask = nMaxMask
		}

		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, ipIndex)
		ipNets = append(ipNets, &net.IPNet{
			IP:   net.IPv4(b[0], b[1], b[2], b[3]),
			Mask: net.CIDRMask(mask, 32),
		})

		// 移动浮标，减数量
		ipIndex += 1 << uint32(32-mask)
		n -= 1 << uint32(32-mask)
	}
	return ipNets
}

// bitLength
// 17: 0001 0001 返回5
// 8: 0000 1000 返回4
func bitLength(x int) int {
	if x == 0 {
		return 1
	}
	var ans = 0
	for x > 0 {
		x >>= 1
		ans++
	}
	return ans
}
