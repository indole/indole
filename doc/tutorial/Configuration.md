# Configuration

INDOLE configuration is in xml format.

# Schema

1. The root node of INDOLE configuration must be `<Indole>`
2. The child nodes of `<Indole>` must be `<Manager>`. Managers can manage plugins, connections and a control to create pipes. For more detail about `Manager`, see the reference.
3. The child nodes of `<Manager>` must be some of the followings: 
   1. `<Plugin>`: an interface or transcoder 
   2. `<Connection>`: a connection from one plugin to another
   3. `<Control>`: a controller for management

Here is a example of INDOLE configuration

```xml
<indole>
    <Manager>
        <Plugin name="OpenFileInterface">
            <FileName>src1.txt</FileName>
        </Plugin>
        <Plugin name="CreateFileInterface">
            <FileName>dst1.txt</FileName>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Control name="BasicControl">
        </Control>
    </Manager>
    <Manager>
        <Plugin name="OpenFileInterface">
            <FileName>src2.txt</FileName>
        </Plugin>
        <Plugin name="CreateFileInterface">
            <FileName>dst2.txt</FileName>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Control name="BasicControl">
        </Control>
    </Manager>
</indole>
```
