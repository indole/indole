package externalprocessinterface

import (
	"encoding/xml"
	"indole/manager"
	"io"
	"os/exec"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	cmd := exec.Command(thisptr.Command, thisptr.Args...)
	cmdin, err := cmd.StdinPipe()
	if err != nil {
		return nil
	}
	cmdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil
	}
	cmd.Start()
	return &ExternalProcessInterface{
		cmdin:  cmdin,
		cmdout: cmdout,
		cmd:    cmd,
	}
}

// Args ...
type Args struct {
	Command string   `xml:"Command"`
	Args    []string `xml:"Args"`
}

func init() {
	manager.PluginRegister["ExternalProcessInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
