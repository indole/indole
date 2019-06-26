package plainstream

import (
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	r, w := io.Pipe()
	return &PlainStream{
		reader: r,
		writer: w,
	}
}

// Args ...
type Args struct {
}
