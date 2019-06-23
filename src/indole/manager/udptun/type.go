package udptun

import (
	"indole/core"
	"indole/plugin/tun"
	"indole/plugin/udp"
)

// UDPTUN ...
type UDPTUN struct {
	udp *udp.Args
	tun *tun.Args
	mtu int
}

// Run ...
func (thisptr *UDPTUN) Run() {
	x := udp.Build(thisptr.udp)
	y := tun.Build(&tun.Args{
		Device: "tun0",
	})
	c := make(chan struct{}, 2)

	go core.Core(x, y, thisptr.mtu, c)
	go core.Core(y, x, thisptr.mtu, c)

	select {
	case <-c:
	}
}
