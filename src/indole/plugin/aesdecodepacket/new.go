package aesdecodepacket

import (
	"encoding/hex"
	"encoding/xml"
	"indole/manager"
	"io"
	"log"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	key, err := hex.DecodeString(thisptr.HexKey)
	if err != nil {
		log.Println("[plugin]", "[aesdecodepacket]", "[Build]", "err:", err)
		return nil
	}
	return &AESDecodePacket{
		queue: make(chan []byte, thisptr.QueueSize),
		key:   key,
	}
}

// Args ...
type Args struct {
	QueueSize int    `xml:"QueueSize"`
	HexKey    string `xml:"HexKey"`
}

func init() {
	manager.PluginRegister["AESDecodePacket"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
