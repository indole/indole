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
		Exec: []string{
			"ifconfig tun0 10.8.0.1/16 mtu 1400 up",
			"iptables -t nat -A POSTROUTING -s 10.8.0.0/16 ! -d 10.8.0.0/16 -j MASQUERADE",
			"iptables -A FORWARD -s 10.8.0.0/16 -m state --state RELATED,ESTABLISHED -j ACCEPT",
			"iptables -A FORWARD -d 10.8.0.0/16 -j ACCEPT",
		},
		Exit: []string{
			"iptables -t nat -D POSTROUTING -s 10.8.0.0/16 ! -d 10.8.0.0/16 -j MASQUERADE",
			"iptables -D FORWARD -s 10.8.0.0/16 -m state --state RELATED,ESTABLISHED -j ACCEPT",
			"iptables -D FORWARD -d 10.8.0.0/16 -j ACCEPT",
		},
	})
	c := make(chan struct{}, 2)

	go core.Core(x, y, thisptr.mtu, c)
	go core.Core(y, x, thisptr.mtu, c)

	select {
	case <-c:
	}
}
