package subnet

type Ipv4Range struct {
	Start uint32
	End   uint32
}

// 用于排序
type Ipv4Ranges []*Ipv4Range

func (i Ipv4Ranges) Len() int {
	return len(i)
}
func (i Ipv4Ranges) Less(j, k int) bool {
	return i[j].Start < i[k].Start
}
func (i Ipv4Ranges) Swap(j, k int) {
	i[j], i[k] = i[k], i[j]
}
