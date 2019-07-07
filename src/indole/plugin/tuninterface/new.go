package tuninterface

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
		// When the device shut down,routing table which was related by interface will delete automatically
		// for _, s := range thisptr.Route {
		// 	croute := C.CString(s)
		// 	delroute := C.del_route_cidr(cdev, croute)
		// 	if delroute < 0 {
		// 		log.Println("plugin", "tuninterface", "delroute", "del route failed", s)
		// 	}
		// 	C.free(unsafe.Pointer(croute))
		// }
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
	for _, s := range thisptr.Route {
		croute := C.CString(s)
		addroute := C.add_route_cidr(cdev, croute)
		if addroute < 0 {
			log.Println("plugin", "tuninterface", "addroute", "add route failed", s)
		}
		C.free(unsafe.Pointer(croute))
	}
	if tunfd < 0 {
		log.Println("plugin", "tuninterface", "New", "set up failed")
		return nil
	}
	if setip < 0 {
		log.Println("plugin", "tuninterface", "set_ip", "set ip failed")
		return nil
	}
	if setmtu < 0 {
		log.Println("plugin", "tuninterface", "set mtu", "set mtu failed")
		return nil
	}
	return &TUNInterface{os.NewFile(uintptr(tunfd), thisptr.Device)}
}

// Args ...
type Args struct {
	Device  string   `xml:"Device"`
	Ipaddr  string   `xml:"Ipaddr"`
	Netmask string   `xml:"Netmask"`
	Mtu     int      `xml:"Mtu"`
	Route   []string `xml:"Route"`
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
