package createfileinterface

import (
	"encoding/xml"
	"indole/manager"
	"io"
	"os"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	if f, err := os.Create(thisptr.FileName); err == nil {
		return &CreateFileInterface{f}
	}
	return nil
}

// Args ...
type Args struct {
	FileName string `xml:"FileName"`
}

func init() {
	manager.PluginRegister["CreateFileInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
