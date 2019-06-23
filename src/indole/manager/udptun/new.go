package udptun

import (
	"indole/manager"
	"indole/plugin/tun"
	"indole/plugin/udp"
)

// Build ...
func Build(args *Args) manager.Manager {
	return &UDPTUN{
		udp: args.UDP,
		tun: args.TUN,
		mtu: args.MTU,
	}
}

// Args ...
type Args struct {
	UDP *udp.Args `xml:"udp"`
	TUN *tun.Args `xml:"tun"`
	MTU int       `xml:"mtu,attr"`
}
