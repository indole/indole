package hexencodestream

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
		encoder: hex.NewEncoder(w),
	}
}

// Args ...
type Args struct {
}
