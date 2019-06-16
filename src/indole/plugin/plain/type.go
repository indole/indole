package plain

// PLAIN ...
type PLAIN struct {
	queue chan []byte
}

// Close ...
func (thisptr *PLAIN) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *PLAIN) Read(p []byte) (n int, err error) {
	data := <-thisptr.queue
	copy(p, data)
	return len(data), nil
}

// Write ...
func (thisptr *PLAIN) Write(p []byte) (n int, err error) {
	n = len(p)
	data := make([]byte, n)
	copy(data, p)
	thisptr.queue <- data
	return
}
