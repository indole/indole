package tuninterface

/*
#include <stdint.h>
int32_t setup_tun_device(char *ifname);
int32_t set_ip(char *name, char *ip_addr, char *netmask);
*/
import "C"
import (
	"encoding/xml"
	"indole/manager"
	"io"
	"log"
	"os"
	"unsafe"
)

//Build ...
func Build(args *Args) io.ReadWriteCloser {
	tunfd := C.setup_tun_device((*C.char)(unsafe.Pointer(&args.Device)))
	setip := C.set_ip(
		(*C.char)(unsafe.Pointer(&args.Device)),
		(*C.char)(unsafe.Pointer(&args.Ipaddr)),
		(*C.char)(unsafe.Pointer(&args.Netmask)),
	)
	if tunfd < 0 || setip < 0 {
		log.Println("plugin", "tuninterface", "New", "C.setup_tun_device()")
		return nil
	}
	return &TUN{os.NewFile(uintptr(tunfd), args.Device)}
}

// Args ...
type Args struct {
	Device  string `xml:"device,attr"`
	Ipaddr  string `xml:"ip,attr"`
	Netmask string `xml:"netmask,attr"`
}

func init() {
	manager.PluginRegister["tuninterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
