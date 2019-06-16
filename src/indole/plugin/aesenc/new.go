package aesenc

import (
	"bytes"
	"encoding/hex"
	"io"
	"log"
)

// New ...
func New(queue chan []byte, key []byte, buffer *bytes.Buffer) io.ReadWriteCloser {
	return &AESENC{
		queue:  queue,
		key:    key,
		buffer: buffer,
	}
}

// NewBySizeHexKey ...
func NewBySizeHexKey(size int, hexkey string) io.ReadWriteCloser {
	key, err := hex.DecodeString(hexkey)
	if err != nil {
		log.Fatalln("plugin", "aesenc", "NewBySizeHexKey", err)
	}
	return New(make(chan []byte, size), key, bytes.NewBuffer(make([]byte, 0)))
}
