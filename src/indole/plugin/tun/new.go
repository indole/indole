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
func New(config map[string]interface{}) io.ReadWriteCloser {
	log.Println("vertex", "tun", "new", "New", config)
	dev := config["dev"].(string)
	cdev := C.CString(dev)
	defer C.free(unsafe.Pointer(cdev))
	tunfd := C.tun_alloc(cdev)
	dev = C.GoString(cdev)
	if tunfd < 0 {
		log.Fatalln("vertex", "tun", "new", "New", "C.tun_alloc(cdev)")
	}

	for _, v := range config["exec"].([]interface{}) {
		err := exec.Command("sh", "-c", v.(string)).Run()
		if err != nil {
			log.Fatalln("vertex", "tun", "new", "New", v, err)
		}
	}

	return &TUN{
		file: os.NewFile(uintptr(tunfd), dev),
		exit: func() {
			for _, v := range config["exit"].([]interface{}) {
				err := exec.Command("sh", "-c", v.(string)).Run()
				if err != nil {
					log.Fatalln("vertex", "tun", "new", "New", v, err)
				}
			}
		},
	}
}
