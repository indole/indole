package tun

import (
	"os"
)

// TUN ...
type TUN struct {
	*os.File
}
