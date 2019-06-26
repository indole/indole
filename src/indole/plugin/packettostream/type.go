package packettostream

import (
	"encoding/binary"
	"indole/utils"
	"io"
)

// PacketToStream ...
type PacketToStream struct {
	reader *io.PipeReader
	writer *io.PipeWriter
}

// Close ...
func (thisptr *PacketToStream) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *PacketToStream) Read(p []byte) (n int, err error) {
	return thisptr.reader.Read(p)
}

// Write ...
func (thisptr *PacketToStream) Write(p []byte) (n int, err error) {
	if err := binary.Write(thisptr.writer, binary.LittleEndian, uint64(len(p))); err != nil {
		return 0, err
	}
	return thisptr.writer.Write(p)
}
