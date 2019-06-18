package aesdec

import (
	"bytes"
	"encoding/hex"
	"io"
	"log"
)

// New ...
func New(args *Args) io.ReadWriteCloser {
	log.Println("plugin", "aesdec", "New", args)
	key, err := hex.DecodeString(args.HexKey)
	if err != nil {
		log.Fatalln("plugin", "aesenc", "NewBySizeHexKey", err)
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
