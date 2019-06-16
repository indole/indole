package tun

import (
	"log"
	"os"
)

// TUN ...
type TUN struct {
	file *os.File
	exit func()
}

// Close ...
func (thisptr *TUN) Close() error {
	defer thisptr.exit()
	log.Println("vertex", "tun", "type", "Close")
	return thisptr.file.Close()
}

// Read ...
func (thisptr *TUN) Read(p []byte) (n int, err error) {
	return thisptr.file.Read(p)
}

// Write ...
func (thisptr *TUN) Write(p []byte) (n int, err error) {
	return thisptr.file.Write(p)
}
