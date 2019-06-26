package plainstream

import (
	"indole/utils"
	"io"
)

// PlainStream ...
type PlainStream struct {
	reader *io.PipeReader
	writer *io.PipeWriter
}

// Close ...
func (thisptr *PlainStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *PlainStream) Read(p []byte) (n int, err error) {
	return thisptr.reader.Read(p)
}

// Write ...
func (thisptr *PlainStream) Write(p []byte) (n int, err error) {
	return thisptr.writer.Write(p)
}
