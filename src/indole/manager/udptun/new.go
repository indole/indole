package udptun

import (
	"indole/manager"
)

// Build ...
func Build(args *Args) manager.Manager {
	return &UDPTUN{}
}

// Args ...
type Args struct {
}
