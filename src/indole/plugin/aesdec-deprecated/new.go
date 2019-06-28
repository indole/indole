package aesdec

import (
	"bytes"
	"encoding/hex"
	"io"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	key, err := hex.DecodeString(args.HexKey)
	if err != nil {
		return nil
	}
	r, w := io.Pipe()
	return &AESDEC{
		queue:  make(chan []byte, args.QueueSize),
		key:    key,
		r:      r,
		w:      w,
		buf:    make([]byte, args.BufSize),
		buffer: bytes.NewBuffer(make([]byte, args.BufferInitSize)),
	}
}

// Args ...
type Args struct {
	QueueSize      int    `xml:"queue_size,attr"`
	HexKey         string `xml:"hex_key,attr"`
	BufSize        int    `xml:"buf_size,attr"`
	BufferInitSize int    `xml:"buffer_init_size,attr"`
}
