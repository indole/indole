# LocalUDPToRemoteUDP

```xml
<Indole>
    <Manager>
        <Plugin name="DynamicUDPInterface">
            <Network>udp</Network>
            <Address>:54345</Address>
        </Plugin>
        <Plugin name="UDPInterface">
            <Network>udp</Network>
            <Address>:54346</Address>
            <RemoteNetwork>udp</RemoteNetwork>
            <RemoteAddress>202.102.3.3:54345</RemoteAddress>
        </Plugin>
        <Connection x="0" y="1" size="1400"/>
        <Connection x="1" y="0" size="1400"/>
        <Control name="BasicControl">
        </Control>
    </Manager>
</Indole>
```