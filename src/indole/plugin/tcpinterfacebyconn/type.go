package tcpinterfacebyconn

import (
	"net"
)

// TCPInterfaceByConn ...
type TCPInterfaceByConn struct {
	net.Conn
}
