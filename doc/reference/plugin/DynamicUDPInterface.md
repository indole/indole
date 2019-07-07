# DynamicUDPInterface

UDP interface that pair the lastest incomming remote address

## Args

| name    | type   | description                                                                                                                                |
| ------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| Network | string | network to listen, which can be `udp`, `udp4` (IPv4-only), `udp6` (IPv6-only), see also [ListenUDP](https://golang.org/pkg/net/#ListenUDP) |
| Address | string | address to listen, see also [ListenUDP](https://golang.org/pkg/net/#ListenUDP)                                                             |


## IO

| --- | type   | size |
| --- | ------ | ---- |
| I   | packet | `-`  |
| O   | packet | `-`  |

## Related

[UDPInterface](UDPInterface.md)
[UDPInterfaceWriteErrorIgnore](UDPInterfaceWriteErrorIgnore.md)
