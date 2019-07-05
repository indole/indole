# DynamicUDPInterface

UDP interface that pair the lastest incomming remote address

## Args

| name    | type   | description                                                                                                                                                 |
| ------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Network | string | network to listen, which can be `udp`, `udp4` (IPv4-only), `udp6` (IPv6-only), see also [ListenUDP](https://golang.org/src/net/udpsock.go?s=6961:7025#L221) |
| Address | string | address to listen, see also [ListenUDP](https://golang.org/src/net/udpsock.go?s=6961:7025#L221)                                                             |


## IO

| --- | type   | size              |
| --- | ------ | ----------------- |
| I   | stream | `-`               |
| O   | stream | `-` |

## Related
