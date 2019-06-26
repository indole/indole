package hexdecodestream

import (
	"indole/utils"
	"io"
)

// HexDecodeStream ...
type HexDecodeStream struct {
	reader  *io.PipeReader
	writer  *io.PipeWriter
	decoder io.Reader
}

// Close ...
func (thisptr *HexDecodeStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *HexDecodeStream) Read(p []byte) (n int, err error) {
	return thisptr.decoder.Read(p)
}

// Write ...
func (thisptr *HexDecodeStream) Write(p []byte) (n int, err error) {
	return thisptr.writer.Write(p)
}
