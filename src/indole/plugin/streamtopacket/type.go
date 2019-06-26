package streamtopacket

import (
	"encoding/binary"
	"indole/utils"
	"io"
)

// StreamToPacket ...
type StreamToPacket struct {
	queue  chan []byte
	reader *io.PipeReader
	writer *io.PipeWriter
}

// Close ...
func (thisptr *StreamToPacket) Close() error {
	close(thisptr.queue)
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *StreamToPacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *StreamToPacket) Write(p []byte) (n int, err error) {
	return thisptr.writer.Write(p)
}

func (thisptr *StreamToPacket) transcode() {
	defer func() {
		recover()
	}()
	for {
		var s uint64
		if err := binary.Read(thisptr.reader, binary.LittleEndian, &s); err != nil {
			thisptr.Close()
			return
		}
		size := int(s)
		buffer := make([]byte, size)
		for i := 0; i < size; {
			n, err := thisptr.reader.Read(buffer[i:size])
			if err != nil {
				thisptr.Close()
				return
			}
			i += n
		}
		thisptr.queue <- buffer
	}
}
