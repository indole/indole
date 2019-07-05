# TCPInterface

connect a tcp address by dialing

## Args

| name    | type   | description                                      |
| ------- | ------ | ------------------------------------------------ |
| Network | string | network to listen, which can be `tcp`, `tcp4` (IPv4-only), `tcp6` (IPv6-only), see also [Listen](https://golang.org/pkg/net/#Listen) |
| Address | string | address to listen, see also [Listen](https://golang.org/pkg/net/#Listen)   


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | packet | `-`    |
| O   | packet | `-` |

## Related

[TCPInterfaceByConn](TCPInterfaceByConn.md)