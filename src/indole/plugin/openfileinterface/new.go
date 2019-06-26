package openfileinterface

import (
	"encoding/xml"
	"indole/manager"
	"io"
	"os"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	if f, err := os.Open(thisptr.FileName); err == nil {
		return &OpenFileInterface{f}
	}
	return nil
}

// Args ...
type Args struct {
	FileName string `xml:"FileName"`
}

func init() {
	manager.PluginRegister["OpenFileInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
