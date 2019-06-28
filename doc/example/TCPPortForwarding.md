# TCPPortForwarding

```xml
<Indole>
    <Manager>
        <Plugin name="TCPInterface">
            <Network>tcp</Network>
            <Address>0.0.0.0:3024</Address>
        </Plugin>
        <Control Name="TCPControl">
            <Network>tcp</Network>
            <Address>0.0.0.0:3025</Address>
            <In>0</In>
            <Out>4</Out>
            <Size>4096</Size>
        </Control>
    </Manager>
</Indole>
```

```mermaid
graph LR
F((Browser))
A(localhost:3025)
B[localhost:3024]
F-.->A
A-->B
B-->A
A-.->F
```