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
| I   | stream | `n`              |
| O   | stream | `[n+8, n+Size+8]` |

## Related
