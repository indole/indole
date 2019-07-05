# AESEncodePacket

AES Encode with `AES-128-CFB` or `AES-256-CFB`

## Args

| name      | type   | description                                      |
| --------- | ------ | ------------------------------------------------ |
| QueueSize | int    | the package queue size _(optional, default `0`)_ |
| HexKey    | string | the secret key in aes decryption in hex format   |


## IO

| --- | type   | size   |
| --- | ------ | ------ |
| I   | packet | `n`    |
| O   | packet | `n+16` |
