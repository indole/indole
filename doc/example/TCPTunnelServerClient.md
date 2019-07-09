# TCPTunnelServerClient

## Server

```xml
<Indole>
    <Manager>
        <Plugin name="StreamToPacketWithAES">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
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
        <Plugin name="PacketToStreamWithAES">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Connection x="1" y="2" size="4096"/>
        <Connection x="2" y="3" size="4096"/>
        <Connection x="3" y="4" size="8192"/>
        <Control name="TCPControl">
            <Network>tcp</Network>
            <Address>0.0.0.0:3023</Address>
            <In>0</In>
            <Out>4</Out>
            <Size>4096</Size>
        </Control>
    </Manager>
</Indole>
```

## Client

```xml
<Indole>
    <Manager>
        <Plugin name="AESEncodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="PacketToStreamWithAES">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="TCPInterface">
            <Network>tcp</Network>
            <Address>server.example.com:3023</Address>
        </Plugin>
        <Plugin name="StreamToPacketWithAES">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="AESDecodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Connection x="1" y="2" size="4096"/>
        <Connection x="2" y="3" size="4096"/>
        <Connection x="3" y="4" size="8192"/>
        <Control name="TCPControl">
            <Network>tcp</Network>
            <Address>localhost:3023</Address>
            <In>0</In>
            <Out>4</Out>
            <Size>4096</Size>
        </Control>
    </Manager>
</Indole>
```
