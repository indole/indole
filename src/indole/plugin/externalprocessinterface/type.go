package externalprocessinterface

import (
	"indole/utils"
	"io"
	"os/exec"
)

// ExternalProcessInterface ...
type ExternalProcessInterface struct {
	cmdin  io.WriteCloser
	cmdout io.ReadCloser
	cmd    *exec.Cmd
}

// Close ...
func (thisptr *ExternalProcessInterface) Close() error {
	return utils.FirstError(thisptr.cmdin.Close(), thisptr.cmdout.Close(), thisptr.cmd.Wait())
}

// Read ...
func (thisptr *ExternalProcessInterface) Read(p []byte) (n int, err error) {
	return thisptr.cmdout.Read(p)
}

// Write ...
func (thisptr *ExternalProcessInterface) Write(p []byte) (n int, err error) {
	return thisptr.cmdin.Write(p)
}
