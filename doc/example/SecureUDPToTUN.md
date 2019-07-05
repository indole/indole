# UDPToTUN

```xml
<Indole>
    <Manager>
        <Plugin name="DynamicUDPInterface">
            <Network>udp</Network>
            <Address>:54345</Address>
        </Plugin>
        <Plugin name="AESDecodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="AESEncodePacket">
            <HexKey>ffffffffffffffffffffffffffffffff</HexKey>
        </Plugin>
        <Plugin name="TUNInterface">
            <Device>tun0</Device>
        </Plugin>
        <Connection x="0" y="1" size="1500"/>
        <Connection x="1" y="3" size="1400"/>
        <Connection x="3" y="2" size="1400"/>
        <Connection x="2" y="0" size="1500"/>
        <Control name="BasicControl">
        </Control>
    </Manager>
</Indole>
```