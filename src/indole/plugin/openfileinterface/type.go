package openfileinterface

import (
	"os"
)

// OpenFileInterface ...
type OpenFileInterface struct {
	*os.File
}
