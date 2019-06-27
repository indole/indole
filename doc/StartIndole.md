# Start INDOLE

The basic way to start INDOLE is:

```sh
./indole
```

Specially, for windows command prompt users:

```cmd
indole.exe
```

And windows powershell users:

```powershell
.\indole.exe
```

# Input Configuration

You can input the configuration xml after you start INDOLE, a configuration example is here:

```xml
<indole>
    <Manager>
        <Plugin Name="OpenFileInterface">
            <FileName>src.txt</FileName>
        </Plugin>
        <Plugin Name="CreateFileInterface">
            <FileName>dst.txt</FileName>
        </Plugin>
        <Connection x="0" y="1" size="8192"/>
        <Control Name="BasicControl">
        </Control>
    </Manager>
</indole>
```

# Send Configuration File to Stdin

You can send the configuration file to pipe while starting INDOLE

```sh
./indole < config.xml
```

Specially, for windows command prompt users:

```cmd
indole.exe < config.xml
```

And windows powershell users:

```powershell
cat config.xml | .\indole.exe
```