# Configuration

INDOLE configuration is in xml format.

# Schema

Here is a example of INDOLE configuration

```xml
<indole>
    <Manager>
        <Plugin Name="OpenFileInterface">
            <FileName>src1.txt</FileName>
        </Plugin>
        <Plugin Name="CreateFileInterface">
            <FileName>dst1.txt</FileName>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Control Name="BasicControl">
        </Control>
    </Manager>
    <Manager>
        <Plugin Name="OpenFileInterface">
            <FileName>src2.txt</FileName>
        </Plugin>
        <Plugin Name="CreateFileInterface">
            <FileName>dst2.txt</FileName>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Control Name="BasicControl">
        </Control>
    </Manager>
</indole>
```

1. The root node of INDOLE configuration must be `<indole>`
2. The child nodes of `<indole>` must be `<Manager>`. Managers can manage plugins, connections and a control to create pipes. For more detail about `Manager`, see the reference.
3. TODO