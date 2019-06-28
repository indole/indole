# INDOLE

INDOLE is a data transfer tool focus on privacy protection on Internet

THE PRINCIPLE of INDOLE is:

    an open framework for data transfer with fully customized protocol by end user.

If you want to use INDOLE logo, build it by [open babel](http://openbabel.org). For example:

```sh
obabel indole.sdf -O indole.png -d --title ""
```

# GUI

**DOC** **TODO**

An Windows GUI for Indole (v0.2) [IndoleVPN](https://github.com/Tommo-L/IndoleVPN)

# Usage

## Requirements

1. [gcc](https://gcc.gnu.org/)
2. [golang](https://golang.org/)

> golang is a temporary decision. Welcome new impls especially `rust`, `scheme`, `java`, `c++`

## Build

```sh
env GOPATH=$(pwd) go build indole
```

For windows (`powershell`) users:

```powershell
$env:GOPATH = (gi .)
go build indole
```

## Run

run `indole` and input the configuration (`xml` format) via `stdin`

```sh
./indole < cfg/config.xml
```

For windows (`powershell`) users:

```powershell
cat cfg\config.xml | .\indole.exe
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

**TIME TO NEW INDOLE !!!!!!!!!!!!!!!!**

Example Here **TODO**

```xml
<Indole>
    <Manager>
        <Plugin name="AESEncodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="PacketToStream">
        </Plugin>
        <Plugin name="TCPInterface">
            <Network>tcp</Network>
            <Address>127.0.0.1:3024</Address>
        </Plugin>
        <Connection x="0" y="1" size="2048"/>
        <Connection x="1" y="2" size="1024"/>
        <Control name="TCPControl">
            <Network>tcp</Network>
            <Address>0.0.0.0:3025</Address>
            <In>0</In>
            <Out>2</Out>
            <Size>1024</Size>
        </Control>
    </Manager>

    <Manager>
        <Plugin name="StreamToPacket">
        </Plugin>
        <Plugin name="AESDecodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="TCPInterface">
            <Network>tcp</Network>
            <Address>127.0.0.1:8118</Address>
        </Plugin>
        <Connection x="0" y="1" size="2048"/>
        <Connection x="1" y="2" size="1024"/>
        <Control name="TCPControl">
            <Network>tcp</Network>
            <Address>0.0.0.0:3024</Address>
            <In>0</In>
            <Out>2</Out>
            <Size>1024</Size>
        </Control>
    </Manager>
</Indole>
```

example 2

```xml
<Indole>
    <Manager>
        <Plugin name="AESEncodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="PacketToStream">
        </Plugin>
        <Plugin name="TCPInterface">
            <Network>tcp</Network>
            <Address>127.0.0.1:3024</Address>
        </Plugin>
        <Plugin name="StreamToPacket">
        </Plugin>
        <Plugin name="AESDecodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Connection x="0" y="1" size="2048"/>
        <Connection x="1" y="2" size="1024"/>
        <Connection x="2" y="3" size="1024"/>
        <Connection x="3" y="4" size="2048"/>
        <Control name="TCPControl">
            <Network>tcp</Network>
            <Address>0.0.0.0:3025</Address>
            <In>0</In>
            <Out>4</Out>
            <Size>1024</Size>
        </Control>
    </Manager>

    <Manager>
        <Plugin name="StreamToPacket">
        </Plugin>
        <Plugin name="AESDecodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="TCPInterface">
            <Network>tcp</Network>
            <Address>127.0.0.1:8118</Address>
        </Plugin>
        <Plugin name="AESEncodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="PacketToStream">
        </Plugin>
        <Connection x="0" y="1" size="2048"/>
        <Connection x="1" y="2" size="1024"/>
        <Connection x="2" y="3" size="1024"/>
        <Connection x="3" y="4" size="2048"/>
        <Control name="TCPControl">
            <Network>tcp</Network>
            <Address>0.0.0.0:3024</Address>
            <In>0</In>
            <Out>4</Out>
            <Size>1024</Size>
        </Control>
    </Manager>
</Indole>
```