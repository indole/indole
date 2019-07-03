package udpinterfacewriteerrorignore

import (
	"encoding/xml"
	"indole/manager"
	"io"
	"net"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	addr, err := net.ResolveUDPAddr(thisptr.Network, thisptr.Address)
	if err != nil {
		return nil
	}
	remote, err := net.ResolveUDPAddr(thisptr.RemoteNetwork, thisptr.RemoteAddress)
	if err != nil {
		return nil
	}
	conn, err := net.ListenUDP(thisptr.Network, addr)
	if err != nil {
		return nil
	}
	return &UDPInterfaceWriteErrorIgnore{
		conn:   conn,
		remote: remote,
	}
}

// Args ...
type Args struct {
	Network       string `xml:"Network"`
	Address       string `xml:"Address"`
	RemoteNetwork string `xml:"RemoteNetwork"`
	RemoteAddress string `xml:"RemoteAddress"`
}

func init() {
	manager.PluginRegister["UDPInterfaceWriteErrorIgnore"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
