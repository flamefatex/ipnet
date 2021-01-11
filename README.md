# ipnet

## subnet

### merger

将连续的网段进行合并，然后输出
```golang
m := subnet.NewSubnetMerger()

err = m.AppendIpNet(ipNet)
if err != nil {
    return err
}

ipNets, err := m.Merge()
if err != nil {
    return err
}
```