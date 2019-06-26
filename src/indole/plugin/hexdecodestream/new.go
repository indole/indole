package hexdecodestream

import (
	"encoding/hex"
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	r, w := io.Pipe()
	return &HexDecodeStream{
		reader:  r,
		writer:  w,
		decoder: hex.NewDecoder(r),
	}
}

// Args ...
type Args struct {
}

func init() {
	manager.PluginRegister["HexDecodeStream"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
