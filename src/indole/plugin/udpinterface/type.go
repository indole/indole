package udpinterface

import (
	"net"
)

// UDPInterface ...
type UDPInterface struct {
	conn   *net.UDPConn
	remote *net.UDPAddr
}

// Close ...
func (thisptr *UDPInterface) Close() error {
	return thisptr.conn.Close()
}

// Read ...
func (thisptr *UDPInterface) Read(p []byte) (n int, err error) {
	for {
		n, addr, err := thisptr.conn.ReadFromUDP(p)
		if err != nil || (addr.IP.Equal(thisptr.remote.IP) && addr.Port == thisptr.remote.Port && addr.Zone == thisptr.remote.Zone) {
			return n, err
		}
	}
}

// Write ...
func (thisptr *UDPInterface) Write(p []byte) (n int, err error) {
	return thisptr.conn.WriteToUDP(p, thisptr.remote)
}
