package streamtopacket

import (
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	r, w := io.Pipe()
	ret := &StreamToPacket{
		queue:  make(chan []byte, thisptr.QueueSize),
		reader: r,
		writer: w,
	}
	go ret.transcode()
	return ret
}

// Args ...
type Args struct {
	QueueSize int `xml:"QueueSize"`
}

func init() {
	manager.PluginRegister["StreamToPacket"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
