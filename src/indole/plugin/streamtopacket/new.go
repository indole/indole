package streamtopacket

import (
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	r, w := io.Pipe()
	ret := &StreamToPacket{
		queue:  make(chan []byte, args.QueueSize),
		reader: r,
		writer: w,
	}
	go ret.transcode()
	return ret
}

// Args ...
type Args struct {
	QueueSize int `xml:"queue_size,attr"`
}
