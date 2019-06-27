package packettostreamwithaes

import (
	"encoding/hex"
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	key, err := hex.DecodeString(thisptr.HexKey)
	if err != nil {
		return nil
	}
	r, w := io.Pipe()
	return &PacketToStreamWithAES{
		reader: r,
		writer: w,
		key:    key,
	}
}

// Args ...
type Args struct {
	HexKey string `xml:"HexKey"`
}

func init() {
	manager.PluginRegister["PacketToStreamWithAES"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
