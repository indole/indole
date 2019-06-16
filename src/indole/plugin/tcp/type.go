package tcp

import (
	"net"
)

// TCP ...
type TCP struct {
	conn net.Conn
}

// Close ...
func (thisptr *TCP) Close() error {
	return thisptr.conn.Close()
}

// Read ...
func (thisptr *TCP) Read(p []byte) (n int, err error) {
	return thisptr.conn.Read(p)
}

// Write ...
func (thisptr *TCP) Write(p []byte) (n int, err error) {
	return thisptr.conn.Write(p)
}
