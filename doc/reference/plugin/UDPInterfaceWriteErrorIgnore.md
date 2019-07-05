# UDPInterfaceWriteErrorIgnore

listen a udp address, writing data to remote address.

The error while writing data is ignored. (A udp sevice may use this plugin to survive in intermittent network failures)

## Args

| name      | type   | description                                      |
| --------- | ------ | ------------------------------------------------ |
| QueueSize | int    | the package queue size _(optional, default `0`)_ |
| HexKey    | string | the secret key in aes decryption in hex format   |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | packet | `-`    |
| O   | packet | `-` |

## Related

[UDPInterface](UDPInterface.md)
[DynamicUDPInterface](DynamicUDPInterface.md)
