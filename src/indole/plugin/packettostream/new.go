package packettostream

import (
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	r, w := io.Pipe()
	return &PacketToStream{
		reader: r,
		writer: w,
	}
}

// Args ...
type Args struct {
}
