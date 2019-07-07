# StreamToPacket

transform stream to packet, detecting the packet length from header

## Args

| name      | type | description                                      |
| --------- | ---- | ------------------------------------------------ |
| QueueSize | int  | the package queue size _(optional, default `0`)_ |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | stream | `n`    |
| O   | packet | `n-8` |

## Related

[PacketToStream](PacketToStream.md)
[PacketToStreamWithAES](PacketToStreamWithAES.md)
[StreamToPacketWithAES](StreamToPacketWithAES.md)