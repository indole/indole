package udp

import (
	"io"
	"log"
	"net"
	"sync/atomic"
)

// New ...
func New(args *Args) io.ReadWriteCloser {
	addr, err := net.ResolveUDPAddr(args.Network, args.Address)
	if err != nil {
		log.Println("plugin", "udp", "New", err)
		return nil
	}

	conn, err := net.ListenUDP(args.Network, addr)
	if err != nil {
		log.Println("vertex", "udp", "new", "New", "net.ListenUDP", err)
	}
	var rmt *net.UDPAddr

	var remote atomic.Value
	ret := &UDP{
		conn:   conn,
		remote: remote,
	}
	ret.remote.Store(rmt)
	return ret
}

// Args ...
type Args struct {
	Network string `xml:"network,attr"`
	Address string `xml:"address,attr"`
}
