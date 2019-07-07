package paddingdecodepacket

import (
	"encoding/binary"
	"errors"
)

// PaddingDecodePacket ...
type PaddingDecodePacket struct {
	queue chan []byte
}

// Close ...
func (thisptr *PaddingDecodePacket) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *PaddingDecodePacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *PaddingDecodePacket) Write(p []byte) (n int, err error) {
	if len(p) < 8 {
		return 0, errnel
	}
	size := int(binary.LittleEndian.Uint64(p[len(p)-8:]))
	if len(p)-8 < size {
		return 0, errnel
	}
	data := make([]byte, size)
	copy(data, p[:size])
	thisptr.queue <- data
	return len(p), nil
}

var errnel = errors.New("not enough length")
