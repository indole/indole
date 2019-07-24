package base64encodestream

import (
	"encoding/base64"
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	r, w := io.Pipe()
	return &Base64EncodeStream{
		reader:  r,
		writer:  w,
		encoder: base64.NewEncoder(base64.StdEncoding, w),
	}
}

// Args ...
type Args struct {
}

func init() {
	manager.PluginRegister["Base64EncodeStream"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
