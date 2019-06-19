package tun

// #include <stdlib.h>
// int tun_alloc(char* dev);
import "C"
import (
	"io"
	"log"
	"os"
	"os/exec"
	"unsafe"
)

// New ...
func New(args *Args) io.ReadWriteCloser {
	cdev := C.CString(args.Device)
	defer C.free(unsafe.Pointer(cdev))
	tunfd := C.tun_alloc(cdev)
	dev := C.GoString(cdev)

	if tunfd < 0 {
		log.Println("plugin", "tun", "New", "C.tun_alloc(cdev)")
		return nil
	}

	for _, v := range args.Exec {
		err := exec.Command("sh", "-c", v).Run()
		if err != nil {
			log.Println("plugin", "tun", "New", v, err)
		}
	}

	return &TUN{
		file: os.NewFile(uintptr(tunfd), dev),
		exit: func() {
			for _, v := range args.Exit {
				err := exec.Command("sh", "-c", v).Run()
				if err != nil {
					log.Println("plugin", "tun", "New", v, err)
				}
			}
		},
	}
}

// Args ...
type Args struct {
	Device string   `xml:"device,attr"`
	Exec   []string `xml:"exec>command"`
	Exit   []string `xml:"exit>command"`
}
