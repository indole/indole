package hexencodestream

import (
	"encoding/hex"
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	r, w := io.Pipe()
	return &HexEncodeStream{
		reader:  r,
		writer:  w,
		encoder: hex.NewEncoder(w),
	}
}

// Args ...
type Args struct {
}

func init() {
	manager.PluginRegister["HexEncodeStream"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
