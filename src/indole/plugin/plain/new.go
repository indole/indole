package plain

import (
	"io"
)

// New ...
func New(args *Args) io.ReadWriteCloser {
	return &PLAIN{
		queue: make(chan []byte, args.QueueSize),
	}
}

// Args ...
type Args struct {
	QueueSize int `xml:"queue_size,attr"`
}
