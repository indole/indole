# TCPTunnel

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