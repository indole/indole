package dynamicpaddingencodepacket

import (
	"encoding/binary"
	"math/rand"
)

// DynamicPaddingEncodePacket ...
type DynamicPaddingEncodePacket struct {
	queue chan []byte
	size  int
}

// Close ...
func (thisptr *DynamicPaddingEncodePacket) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *DynamicPaddingEncodePacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *DynamicPaddingEncodePacket) Write(p []byte) (n int, err error) {
	size := rand.Intn(thisptr.size)
	data := make([]byte, thisptr.size+size+8)
	n = copy(data[:thisptr.size], p)
	binary.LittleEndian.PutUint64(data[thisptr.size+size:], uint64(n))
	thisptr.queue <- data
	return
}
