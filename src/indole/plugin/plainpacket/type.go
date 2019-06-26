package plainpacket

// PlainPacket ...
type PlainPacket struct {
	queue chan []byte
}

// Close ...
func (thisptr *PlainPacket) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *PlainPacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *PlainPacket) Write(p []byte) (n int, err error) {
	data := make([]byte, len(p))
	n = copy(data, p)
	thisptr.queue <- data
	return
}
