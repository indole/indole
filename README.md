# INDOLE

INDOLE is a data transfer tool focus on privacy protection on Internet

THE PRINCIPLE of INDOLE is:

    an open framework for data transfer with fully customized protocol by end user.

# Usage

## Requirements

1. [gcc](https://gcc.gnu.org/)
2. [golang](https://golang.org/)

> golang is a temporary decision. Welcome new impls especially `rust`, `scheme`, `java`, `c++`

## Build

```sh
env GOPATH=$(pwd) go build indole
```

## Run

run `indole` and input the configuration (`xml` format) via `stdin`

```sh
./indole < cfg/config.xml
```

## Deploy

The following tools are recommended for deploy

1. [supervisor](http://supervisord.org/)
2. [docker](https://www.docker.com/)

# Tutotial

start INDOLE

```sh
./indole
```

Then input the xml configuration.

The following of this section will show how the configuration works

## INDOLE Configure Schema

The root node of INDOLE config is `indole`.

```xml
<indole>
</indole>
```

The child nodes of `indole` node should be a manager config. Managers manage the plugins, and organize them to work properly.

## TCPAES Manager

TCPAES manager create a tcp tunnel for data transfer.

The attributes of  `tcpaes` node are:

1. `network`: The network to listen. Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only), "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4" (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and "unixpacket". But only `tcp*` network works because this manager works on tcp.
2. `address`: The address to listen.
3. `bufsize`: The read/write bufsize for core using. If it is too small, the performance would be bad. If it is too large, the memory would be uesd a lot, even crash. It is depend on your QPS and system memory.

The child notes of `tcpaes` are:

1. one `encode` node with some of the following child nodes
   1. `aesenc`: aes encrypt with the following attrbutes
      1. `queue_size`: the buf queue size
      2. `hex_key`: the aes key hex encoded
   2. `aesdec`: aes decrypt with the following attrbutes
      1. `queue_size`: the buf queue size
      2. `hex_key`: the aes key hex encoded
      3. `buf_size`: the buf_size, it is recommend a little larger than the `bufsize` of `tcpaes` depended on the encode method you use.
   3. `plain`: only copy no encrypt or decrypt
2. one `decode` node with child notes same as `encode` node
3. one `tcp` node with the following attributes
   1. `network`: the network to redirect
   2. `address`: the address to redirect

```xml
<indole>
    <tcpaes network="tcp" address="0.0.0.0:3023" bufsize="2048">
        <encode>
            <aesenc queue_size="1024" hex_key="ffffffffffffffffffffffffffffffff"/>
        </encode>
        <decode>
            <aesdec queue_size="1024" hex_key="ffffffffffffffffffffffffffffffff" buf_size="4096"/>
        </decode>
        <tcp network="tcp" address="localhost:8118"/>
    </tcpaes>
</indole>
```

## Server / Client

Here is an example for server and client data transfer.

Note that the `bufsize` and `buf_size` should be optimized

### ServerSide Config

```xml
<indole>
    <tcpaes network="tcp" address="0.0.0.0:<PORT>" bufsize="1024">
        <encode>
            <aesdec queue_size="1024" hex_key="<YOUR AES KEY>" buf_size="65536"/>
        </encode>
        <decode>
            <aesenc queue_size="1024" hex_key="<YOUR AES KEY>"/>
        </decode>
        <tcp network="tcp" address="<YOUR LOCAL ADDRESS AND PORT ON SERVER>"/>
    </tcpaes>
</indole>
```

### ClientSide Config

```xml
<indole>
    <tcpaes network="tcp" address="<YOUR LOCAL ADDRESS AND PORT>" bufsize="1024">
        <encode>
            <aesenc queue_size="1024" hex_key="<YOUR AES KEY>"/>
        </encode>
        <decode>
            <aesdec queue_size="1024" hex_key="<YOUR AES KEY>" buf_size="65536"/>
        </decode>
        <tcp network="tcp" address="<THE SERVER ADDRESS AND PORT>"/>
    </tcpaes>
</indole>
```
