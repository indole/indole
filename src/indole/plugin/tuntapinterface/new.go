package tuntapinterface

/*
#include <stdint.h>
#include <stdlib.h>
int32_t setup_tun_device(char *ifname);
int32_t set_ip(char *name, char *ip_addr, char *netmask);
int32_t set_mtu(char *name,int32_t mtu);
int32_t add_route_cidr(char *name, char *cidr);
int32_t del_route_cidr(char *name, char *cidr);

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
	if tunfd < 0 {
		log.Println("plugin", "tuntapinterface", "New", "set up failed")
		return nil
	}
	if setip < 0 {
		log.Println("plugin", "tuntapinterface", "set_ip", "set ip failed")
		return nil
	}
	if setmtu < 0 {
		log.Println("plugin", "tuntapinterface", "set mtu", "set mtu failed")
		return nil
	}
	return &TUNTAPInterface{os.NewFile(uintptr(tunfd), thisptr.Device)}
}

// Args ...
type Args struct {
	Device  string `xml:"Device"`
	Ipaddr  string `xml:"Ipaddr"`
	Netmask string `xml:"Netmask"`
	Mtu     int    `xml:"Mtu"`
}

func init() {
	manager.PluginRegister["TUNTAPInterface"] = func(config []byte) func() io.ReadWriteCloser {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func() io.ReadWriteCloser {
				return nil
			}
		}
		return args.Build
	}
}
