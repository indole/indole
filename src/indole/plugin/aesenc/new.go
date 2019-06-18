package aesenc

import (
	"bytes"
	"encoding/hex"
	"io"
	"log"
)

// New ...
func New(args *Args) io.ReadWriteCloser {
	log.Println("plugin", "aesenc", "New", args)
	key, err := hex.DecodeString(args.HexKey)
	if err != nil {
		log.Println("plugin", "aesenc", "New", err)
		return nil
	}
	return &AESENC{
		queue:  make(chan []byte, args.QueueSize),
		key:    key,
		buffer: bytes.NewBuffer(make([]byte, args.BufferInitSize)),
	}
}

// Args ...
type Args struct {
	QueueSize      int    `xml:"queue_size,attr"`
	HexKey         string `xml:"hex_key,attr"`
	BufferInitSize int    `xml:"buffer_init_size,attr"`
}
