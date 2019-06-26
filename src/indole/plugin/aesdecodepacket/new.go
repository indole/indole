package aesdecodepacket

import (
	"encoding/hex"
	"io"
	"log"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	key, err := hex.DecodeString(args.HexKey)
	if err != nil {
		log.Println("[plugin]", "[aesdecodepacket]", "[Build]", "err:", err)
		return nil
	}
	return &AESDecodePacket{
		queue: make(chan []byte, args.QueueSize),
		key:   key,
	}
}

// Args ...
type Args struct {
	QueueSize int    `xml:"queue_size,attr"`
	HexKey    string `xml:"hex_key,attr"`
}
