package dynamicudpinterface

import (
	"net"
	"sync"
)

// DynamicUDPInterface ...
type DynamicUDPInterface struct {
	conn   *net.UDPConn
	remote *net.UDPAddr
	lock   sync.Mutex
}

// Close ...
func (thisptr *DynamicUDPInterface) Close() error {
	return thisptr.conn.Close()
}

// Read ...
func (thisptr *DynamicUDPInterface) Read(p []byte) (n int, err error) {
	n, addr, err := thisptr.conn.ReadFromUDP(p)
	thisptr.set(addr)
	return n, err
}

// Write ...
func (thisptr *DynamicUDPInterface) Write(p []byte) (n int, err error) {
	if addr := thisptr.get(); addr != nil {
		return thisptr.conn.WriteToUDP(p, addr)
	}
	return 0, nil
}

func (thisptr *DynamicUDPInterface) set(addr *net.UDPAddr) {
	thisptr.lock.Lock()
	defer thisptr.lock.Unlock()
	thisptr.remote = addr
}

func (thisptr *DynamicUDPInterface) get() *net.UDPAddr {
	thisptr.lock.Lock()
	defer thisptr.lock.Unlock()
	return thisptr.remote
}
