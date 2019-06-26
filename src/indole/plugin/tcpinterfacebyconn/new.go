package tcpinterfacebyconn

import (
	"io"
	"net"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	return &TCPInterfaceByConn{thisptr.Conn}
}

// Args ...
type Args struct {
	Conn net.Conn
}

func init() {
}
