# DynamicPaddingEncodePacket

Padding a packet with dybamic size of zero

## Args

| name      | type | description                                      |
| --------- | ---- | ------------------------------------------------ |
| QueueSize | int  | the package queue size _(optional, default `0`)_ |
| Size      | int  | maximum padding size                             |


## IO

| --- | type   | size             |
| --- | ------ | ---------------- |
| I   | packet | `n`              |
| O   | packet | `[n+8, n+Size+8]` |

## Related

[StaticPaddingEncodePacket](StaticPaddingEncodePacket.md)
[PaddingDecodePacket](PaddingDecodePacket.md)
