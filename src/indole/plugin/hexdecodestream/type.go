package hexdecodestream

import (
	"indole/utils"
	"io"
)

// HexEncodeStream ...
type HexEncodeStream struct {
	reader  *io.PipeReader
	writer  *io.PipeWriter
	decoder io.Reader
}

// Close ...
func (thisptr *HexEncodeStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *HexEncodeStream) Read(p []byte) (n int, err error) {
	return thisptr.decoder.Read(p)
}

// Write ...
func (thisptr *HexEncodeStream) Write(p []byte) (n int, err error) {
	return thisptr.writer.Write(p)
}
