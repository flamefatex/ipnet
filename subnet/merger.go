package subnet

import (
	"encoding/binary"
	"fmt"
	"net"
	"sort"
)

type merger struct {
	ipv4Ranges Ipv4Ranges
}

func NewSubnetMerger() *merger {
	return &merger{
		ipv4Ranges: make([]*Ipv4Range, 0),
	}
}

func (m *merger) AppendIpNet(ipNet *net.IPNet) (err error) {
	// v4
	if ipNet.IP.To4() != nil {
		start := binary.BigEndian.Uint32(ipNet.IP.To4())
		broadcastIp := GetBroadcastIp(ipNet)

		end := binary.BigEndian.Uint32(broadcastIp.To4())

		ipv4Range := &Ipv4Range{
			Start: start,
			End:   end,
		}

		m.ipv4Ranges = append(m.ipv4Ranges, ipv4Range)
	} else {
		err = fmt.Errorf("非ipv4地址")
	}
	return
}

func (m *merger) Merge() (ipNets []*net.IPNet, err error) {
	ipNets = make([]*net.IPNet, 0)

	// 排序
	sort.Sort(m.ipv4Ranges)
	// 循环
	// ipv4 求出一段连续的 start 与 end
	var startIpv4Range, endIpv4Range *Ipv4Range
	for i, ipv4Range := range m.ipv4Ranges {

		if startIpv4Range == nil {
			startIpv4Range = ipv4Range
		}

		// 判断是否有下一个
		nextIndex := i + 1
		if nextIndex < len(m.ipv4Ranges) {
			// 与下一个网段是衔接的
			if ipv4Range.End+1 == m.ipv4Ranges[nextIndex].Start {
				endIpv4Range = m.ipv4Ranges[nextIndex]
				continue
			}
		}

		// 最后一个或者是下一个不衔接
		endIpv4Range = ipv4Range
		n := int(endIpv4Range.End - startIpv4Range.Start + 1)
		ss := Ipv4AndNumToCIDR(startIpv4Range.Start, n)

		for _, s := range ss {
			ipNets = append(ipNets, s)
		}

		// 至空
		startIpv4Range = nil
	}

	return
}
