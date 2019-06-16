package plain

import (
	"io"
)

// New ...
func New(queue chan []byte) io.ReadWriteCloser {
	return &PLAIN{
		queue: queue,
	}
}

// NewBySize ...
func NewBySize(size int) io.ReadWriteCloser {
	return New(make(chan []byte, size))
}
