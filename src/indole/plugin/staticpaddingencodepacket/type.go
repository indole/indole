package staticpaddingencodepacket

import "encoding/binary"

// StaticPaddingEncodePacket ...
type StaticPaddingEncodePacket struct {
	queue chan []byte
	size  int
}

// Close ...
func (thisptr *StaticPaddingEncodePacket) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *StaticPaddingEncodePacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *StaticPaddingEncodePacket) Write(p []byte) (n int, err error) {
	data := make([]byte, thisptr.size+8)
	n = copy(data[:thisptr.size], p)
	binary.LittleEndian.PutUint64(data[thisptr.size:], uint64(n))
	thisptr.queue <- data
	return
}
