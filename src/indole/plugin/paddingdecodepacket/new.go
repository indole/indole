package paddingdecodepacket

import (
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	return &PaddingDecodePacket{
		queue: make(chan []byte, thisptr.QueueSize),
	}
}

// Args ...
type Args struct {
	QueueSize int `xml:"QueueSize"`
}

func init() {
	manager.PluginRegister["PaddingDecodePacket"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
