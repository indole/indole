package tcpinterface

import (
	"encoding/xml"
	"indole/manager"
	"io"
	"net"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	if conn, err := net.Dial(thisptr.Network, thisptr.Address); err == nil {
		return &TCPInterface{conn}
	}
	return nil
}

// Args ...
type Args struct {
	Network string `xml:"Network"`
	Address string `xml:"Address"`
}

func init() {
	manager.PluginRegister["TCPInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
