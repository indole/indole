package plain

import (
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	return &PLAIN{
		queue: make(chan []byte, args.QueueSize),
	}
}

// Args ...
type Args struct {
	QueueSize int `xml:"queue_size,attr"`
}
