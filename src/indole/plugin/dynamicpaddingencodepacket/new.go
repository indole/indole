package dynamicpaddingencodepacket

import (
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	return &DynamicPaddingEncodePacket{
		queue: make(chan []byte, thisptr.QueueSize),
		size:  thisptr.Size,
	}
}

// Args ...
type Args struct {
	QueueSize int `xml:"QueueSize"`
	Size      int `xml:"Size"`
}

func init() {
	manager.PluginRegister["DynamicPaddingEncodePacket"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
