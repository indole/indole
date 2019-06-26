package plainstream

import (
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	r, w := io.Pipe()
	return &PlainStream{
		reader: r,
		writer: w,
	}
}

// Args ...
type Args struct {
}

func init() {
	manager.PluginRegister["PlainStream"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
