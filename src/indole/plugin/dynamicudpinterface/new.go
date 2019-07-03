package dynamicudpinterface

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
	conn, err := net.ListenUDP(thisptr.Network, addr)
	if err != nil {
		return nil
	}
	return &DynamicUDPInterface{
		conn: conn,
	}
}

// Args ...
type Args struct {
	Network string `xml:"Network"`
	Address string `xml:"Address"`
}

func init() {
	manager.PluginRegister["DynamicUDPInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
