package base64encodestream

import (
	"indole/utils"
	"io"
)

// Base64EncodeStream ...
type Base64EncodeStream struct {
	reader  *io.PipeReader
	writer  *io.PipeWriter
	encoder io.Writer
}

// Close ...
func (thisptr *Base64EncodeStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *Base64EncodeStream) Read(p []byte) (n int, err error) {
	return thisptr.reader.Read(p)
}

// Write ...
func (thisptr *Base64EncodeStream) Write(p []byte) (n int, err error) {
	return thisptr.encoder.Write(p)
}
