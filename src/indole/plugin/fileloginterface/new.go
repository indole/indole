package fileloginterface

import (
	"encoding/xml"
	"indole/manager"
	"io"
	"os"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	f, err := os.Create(thisptr.FileName)
	if err != nil {
		return nil
	}
	r, w := io.Pipe()
	return &FileLogInterface{
		reader: r,
		writer: w,
		file:   f,
	}
}

// Args ...
type Args struct {
	FileName string `xml:"FileName"`
}

func init() {
	manager.PluginRegister["FileLogInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
