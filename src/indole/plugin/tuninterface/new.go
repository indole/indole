package tuninterface

/*
#include <stdint.h>
#include <stdlib.h>
int32_t setup_tun_device(char *ifname);
int32_t set_ip(char *name, char *ip_addr, char *netmask);
int32_t set_mtu(char *name,int32_t mtu);
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
func (thisptr *Args) Build() io.ReadWriteCloser {
	cdev := C.CString(thisptr.Device)
	cip := C.CString(thisptr.Ipaddr)
	cnetmask := C.CString(thisptr.Netmask)
	cmtu := C.int(thisptr.Mtu)
	defer func() {
		C.free(unsafe.Pointer(cdev))
		C.free(unsafe.Pointer(cip))
		C.free(unsafe.Pointer(cnetmask))
	}()
	tunfd := C.setup_tun_device(cdev)
	setip := C.set_ip(
		cdev,
		cip,
		cnetmask,
	)
	setmtu := C.set_mtu(
		cdev,
		cmtu,
	)
	if tunfd < 0 || setip < 0 || setmtu < 0 {
		log.Println("plugin", "tuninterface", "New", "set up tundevice")
		return nil
	}
	return &TUNInterface{os.NewFile(uintptr(tunfd), thisptr.Device)}
}

// Args ...
type Args struct {
	Device  string `xml:"Device"`
	Ipaddr  string `xml:"Ipaddr"`
	Netmask string `xml:"Netmask"`
	Mtu     int    `xml:"Mtu"`
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
