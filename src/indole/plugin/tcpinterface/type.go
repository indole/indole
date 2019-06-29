package tcpinterface

import (
	"net"
)

// TCPInterface ...
type TCPInterface struct {
	net.Conn
}
