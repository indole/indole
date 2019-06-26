package hexencodestream

import (
	"indole/utils"
	"io"
)

// HexEncodeStream ...
type HexEncodeStream struct {
	reader  *io.PipeReader
	writer  *io.PipeWriter
	encoder io.Writer
}

// Close ...
func (thisptr *HexEncodeStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *HexEncodeStream) Read(p []byte) (n int, err error) {
	return thisptr.reader.Read(p)
}

// Write ...
func (thisptr *HexEncodeStream) Write(p []byte) (n int, err error) {
	return thisptr.encoder.Write(p)
}
