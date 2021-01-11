# ipnet

## subnet

### merger


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