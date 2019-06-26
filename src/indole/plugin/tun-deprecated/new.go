package tun

// #include <stdlib.h>
// int tun_alloc(char* dev);
import "C"
import (
	"io"
	"log"
	"os"
	"unsafe"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	cdev := C.CString(args.Device)
	defer C.free(unsafe.Pointer(cdev))
	tunfd := C.tun_alloc(cdev)
	dev := C.GoString(cdev)

	if tunfd < 0 {
		log.Println("plugin", "tun", "New", "C.tun_alloc(cdev)")
		return nil
	}

	return &TUN{os.NewFile(uintptr(tunfd), dev)}
}

// Args ...
type Args struct {
	Device string `xml:"device,attr"`
}
