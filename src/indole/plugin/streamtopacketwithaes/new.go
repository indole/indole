package streamtopacketwithaes

import (
	"encoding/hex"
	"encoding/xml"
	"indole/manager"
	"indole/utils"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	key, err := hex.DecodeString(thisptr.HexKey)
	if err != nil {
		return nil
	}
	r, w := io.Pipe()
	ret := &StreamToPacketWithAES{
		queue:  make(chan []byte, thisptr.QueueSize),
		reader: r,
		writer: w,
		key:    key,
	}
	go func() {
		defer utils.Recover("[WARN]", "[plugin]", "[StreamToPacketWithAES]", "[Build]")
		ret.transcode()
	}()
	return ret
}

// Args ...
type Args struct {
	QueueSize int    `xml:"QueueSize"`
	HexKey    string `xml:"HexKey"`
}

func init() {
	manager.PluginRegister["StreamToPacketWithAES"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
