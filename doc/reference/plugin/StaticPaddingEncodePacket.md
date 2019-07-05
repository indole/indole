# StaticPaddingEncodePacket

Padding a packet to fixed size of zero

## Args

| name      | type | description                                      |
| --------- | ---- | ------------------------------------------------ |
| QueueSize | int  | the package queue size _(optional, default `0`)_ |
| Size      | int  | the fixed packet size                            |


## IO

| --- | type   | size     |
| --- | ------ | -------- |
| I   | packet | `n`      |
| O   | packet | `Size+8` |

## Related

[DynamicPaddingEncodePacket](DynamicPaddingEncodePacket.md)
[PaddingDecodePacket](PaddingDecodePacket.md)
