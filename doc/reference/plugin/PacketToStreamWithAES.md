# PacketToStreamWithAES

transfrom packet to stream with the packet length as header, the header is encoded by `AES-128-CFB` or `AES-256-CFB`

## Args

| name   | type   | description                                    |
| ------ | ------ | ---------------------------------------------- |
| HexKey | string | the secret key in aes decryption in hex format |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | packet | `n`    |
| O   | stream | `n+24` |

## Related

[PacketToStream](PacketToStream.md)
[StreamToPacket](StreamToPacket.md)
[StreamToPacketWithAES](StreamToPacketWithAES.md)
