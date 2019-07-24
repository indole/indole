package base64decodestream

import (
	"indole/utils"
	"io"
)

// Base64DecodeStream ...
type Base64DecodeStream struct {
	reader  *io.PipeReader
	writer  *io.PipeWriter
	decoder io.Reader
}

// Close ...
func (thisptr *Base64DecodeStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *Base64DecodeStream) Read(p []byte) (n int, err error) {
	return thisptr.decoder.Read(p)
}

// Write ...
func (thisptr *Base64DecodeStream) Write(p []byte) (n int, err error) {
	return thisptr.writer.Write(p)
}
