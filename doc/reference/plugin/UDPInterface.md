# UDPInterface

listen a udp address, writing data to remote address

## Args

| name          | type   | description                                                                                                                                |
| ------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| Network       | string | network to listen, which can be `udp`, `udp4` (IPv4-only), `udp6` (IPv6-only), see also [ListenUDP](https://golang.org/pkg/net/#ListenUDP) |
| Address       | string | address to listen, see also [ListenUDP](https://golang.org/pkg/net/#ListenUDP)                                                             |
| RemoteNetwork | string | network to send to, which can be `udp`, `udp4` (IPv4-only), `udp6` (IPv6-only)                                                             |
| RemoteAddress | string | address to send to                                                                                                                         |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | packet | `-`    |
| O   | packet | `-` |

## Related

[DynamicUDPInterface](DynamicUDPInterface.md)
[UDPInterfaceWriteErrorIgnore](UDPInterfaceWriteErrorIgnore.md)
