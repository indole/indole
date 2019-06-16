package tcp

import (
	"io"
	"log"
	"net"
)

// New ...
func New(conn net.Conn) io.ReadWriteCloser {
	return &TCP{
		conn: conn,
	}
}

// NewByDial ...
func NewByDial(netowrk, address string) io.ReadWriteCloser {
	conn, err := net.Dial(netowrk, address)
	if err != nil {
		log.Fatalln("vertex", "tcp", "NewByDial", "net.Dial(netowrk, address)", err)
	}
	return New(conn)
}
