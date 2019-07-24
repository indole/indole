package base64decodestream

import (
	"encoding/base64"
	"encoding/xml"
	"indole/manager"
	"io"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	r, w := io.Pipe()
	return &Base64DecodeStream{
		reader:  r,
		writer:  w,
		decoder: base64.NewDecoder(base64.StdEncoding, r),
	}
}

// Args ...
type Args struct {
}

func init() {
	manager.PluginRegister["Base64DecodeStream"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
