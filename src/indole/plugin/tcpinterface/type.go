package tcpinterface

import (
	"net"
)

// TCPInterface ...
type TCPInterface struct {
	net.Conn
}

func (thisptr *TCPInterface) Write(p []byte) (int, error) {
	return thisptr.Conn.Write(p)
}
