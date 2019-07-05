# PaddingDecodePacket

Remove the padding of packet

## Args

| name      | type | description                                      |
| --------- | ---- | ------------------------------------------------ |
| QueueSize | int  | the package queue size _(optional, default `0`)_ |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | packet | `n`    |
| O   | packet | `[0, n-8]` |

## Related

[DynamicPaddingEncodePacket](DynamicPaddingEncodePacket.md)
[StaticPaddingEncodePacket](StaticPaddingEncodePacket.md)
