package udp

import (
	"net"
	"sync/atomic"
)

// UDP ...
type UDP struct {
	conn   *net.UDPConn
	remote atomic.Value
}

// Close ...
func (thisptr *UDP) Close() error {
	return thisptr.conn.Close()
}

// Read ...
func (thisptr *UDP) Read(p []byte) (n int, err error) {
	n, addr, err := thisptr.conn.ReadFromUDP(p)
	thisptr.remote.Store(addr)
	return n, err
}

// Write ...
func (thisptr *UDP) Write(p []byte) (n int, err error) {
	addr, ok := thisptr.remote.Load().(*net.UDPAddr)
	if ok && addr != nil {
		return thisptr.conn.WriteToUDP(p, addr)
	}
	return 0, nil
}
