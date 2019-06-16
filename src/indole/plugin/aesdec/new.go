package aesdec

import (
	"bytes"
	"encoding/hex"
	"io"
	"log"
)

// New ...
func New(queue chan []byte, key []byte, r *io.PipeReader, w *io.PipeWriter, limit uint64, buffer []byte, bytesbuffer *bytes.Buffer) io.ReadWriteCloser {
	return &AESDEC{
		queue:       queue,
		key:         key,
		r:           r,
		w:           w,
		limit:       limit,
		buffer:      buffer,
		bytesbuffer: bytesbuffer,
	}
}

// NewBySizeHexKeyLimit ...
func NewBySizeHexKeyLimit(size int, hexkey string, limit uint64) io.ReadWriteCloser {
	key, err := hex.DecodeString(hexkey)
	if err != nil {
		log.Fatalln("plugin", "aesenc", "NewBySizeHexKey", err)
	}
	r, w := io.Pipe()
	return New(make(chan []byte, size), key, r, w, limit, make([]byte, limit), bytes.NewBuffer(make([]byte, 0)))
}
