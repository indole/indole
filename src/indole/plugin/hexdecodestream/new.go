package hexdecodestream

import (
	"encoding/hex"
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	r, w := io.Pipe()
	return &HexEncodeStream{
		reader:  r,
		writer:  w,
		decoder: hex.NewDecoder(r),
	}
}

// Args ...
type Args struct {
}
