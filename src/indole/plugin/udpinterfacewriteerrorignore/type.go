package udpinterfacewriteerrorignore

import (
	"net"
)

// UDPInterfaceWriteErrorIgnore ...
type UDPInterfaceWriteErrorIgnore struct {
	conn   *net.UDPConn
	remote *net.UDPAddr
}

// Close ...
func (thisptr *UDPInterfaceWriteErrorIgnore) Close() error {
	return thisptr.conn.Close()
}

// Read ...
func (thisptr *UDPInterfaceWriteErrorIgnore) Read(p []byte) (n int, err error) {
	for {
		n, addr, err := thisptr.conn.ReadFromUDP(p)
		if err != nil || (addr.IP.Equal(thisptr.remote.IP) && addr.Port == thisptr.remote.Port && addr.Zone == thisptr.remote.Zone) {
			return n, err
		}
	}
}

// Write ...
func (thisptr *UDPInterfaceWriteErrorIgnore) Write(p []byte) (n int, err error) {
	thisptr.conn.WriteToUDP(p, thisptr.remote)
	return len(p), nil
}
