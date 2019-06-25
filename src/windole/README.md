# Windole

A GUI for indole for windows

# Build

In this Dir:

```
go get -u github.com/lxn/walk
go get -u github.com/akavel/rsrc
rsrc.exe -manifest windole.exe.manifest -ico indole.ico -o rsrc.syso
go build -ldflags="-H windowsgui"
```

and then follow the readme of indole to build indole:

specially use this commnd to build indole

```
go build indole -ldflags="-H windowsgui"
```