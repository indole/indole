package tuninterface

// #include <stdlib.h>
// int tun_alloc(char* dev);
import "C"
import (
	"encoding/xml"
	"indole/manager"
	"io"
	"os"
	"unsafe"
)

// Build ...
func (thisptr *Args) Build() io.ReadWriteCloser {
	cdev := C.CString(thisptr.Device)
	defer C.free(unsafe.Pointer(cdev))
	f := C.tun_alloc(cdev)
	return &TUNInterface{os.NewFile(uintptr(f), thisptr.Device)}
}

// Args ...
type Args struct {
	Device string `xml:"Device"`
}

func init() {
	manager.PluginRegister["TUNInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
