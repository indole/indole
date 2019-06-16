package udp

import (
	"io"
	"log"
	"net"
	"sync/atomic"
)

// New ...
func New(config map[string]interface{}) io.ReadWriteCloser {
	addr, err := net.ResolveUDPAddr(config["network"].(string), config["address"].(string))
	if err != nil {
		log.Println("vertex", "udp", "new", "New", "net.ResolveUDPAddr", err)
	}

	conn, err := net.ListenUDP(config["network"].(string), addr)
	if err != nil {
		log.Println("vertex", "udp", "new", "New", "net.ListenUDP", err)
	}
	var rmt *net.UDPAddr
	if config["remote"] != nil {
		addr, err := net.ResolveUDPAddr("udp", config["remote"].(string))
		if err != nil {
			log.Println("vertex", "udp", "new", "New", "net.ResolveUDPAddr", err)
		}
		rmt = addr
	}
	var remote atomic.Value
	ret := &UDP{
		conn:   conn,
		remote: remote,
	}
	ret.remote.Store(rmt)
	return ret
}
