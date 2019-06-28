package tuninterface

import (
	"os"
)

// TUN ...
type TUN struct {
	*os.File
}

// func (thisptr *TUN) Write(p []byte) (int, error) {
// 	return thisptr.Conn.Write(p)
// }
