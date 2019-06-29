package tuninterface

import "os"

// TUNInterface ...
type TUNInterface struct {
	*os.File
}
