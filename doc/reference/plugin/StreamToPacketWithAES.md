# StreamToPacketWithAES

transform stream to packet, detecting the packet length from aes encoded header

## Args

| name      | type   | description                                      |
| --------- | ------ | ------------------------------------------------ |
| QueueSize | int    | the package queue size _(optional, default `0`)_ |
| HexKey    | string | the secret key in aes decryption in hex format   |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | stream | `n`    |
| O   | packet | `n-24` |

## Related

[PacketToStream](PacketToStream.md)
[PacketToStreamWithAES](PacketToStreamWithAES.md)
[StreamToPacket](StreamToPacket.md)
