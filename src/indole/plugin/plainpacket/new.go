package plainpacket

import (
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	return &PlainPacket{
		queue: make(chan []byte, args.QueueSize),
	}
}

// Args ...
type Args struct {
	QueueSize int `xml:"queue_size,attr"`
}
