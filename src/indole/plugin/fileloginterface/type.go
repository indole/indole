package fileloginterface

import (
	"indole/utils"
	"io"
	"os"
)

// FileLogInterface ...
type FileLogInterface struct {
	reader *io.PipeReader
	writer *io.PipeWriter
	file   *os.File
}

// Close ...
func (thisptr *FileLogInterface) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close(), thisptr.file.Close())
}

// Read ...
func (thisptr *FileLogInterface) Read(p []byte) (n int, err error) {
	return thisptr.reader.Read(p)
}

// Write ...
func (thisptr *FileLogInterface) Write(p []byte) (n int, err error) {
	if n, err = thisptr.file.Write(p); err != nil {
		return
	}
	return thisptr.writer.Write(p)
}
