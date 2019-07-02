# INDOLE

INDOLE is an open framework for data transfer with fully customized protocol by end user.

# Logo

Here is the sdf file of indole [LOGO](indole.sdf).

If you want another format, build it by [open babel](http://openbabel.org).

```sh
obabel indole.sdf -O indole.png -d --title ""
```

# GUI

- [Windole](src/windole) is a windows GUI for INDOLE.
- [IndoleVPN](https://github.com/Tommo-L/IndoleVPN) is an windows GUI for proxy usage of INDOLE (v0.2) 
- [Droidindole](https://github.com/AaronGarbut/Droidindole) an Android GUI/Interface

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

# Documentation

For more documentation, see [doc](doc)

# Donation

- BTC: bc1q5yvtyhn7uf60xwkck55dl5c909ywae75u8235r