package basiccontrol

import (
	"encoding/xml"
	"indole/manager"
)

// Run ...
func (thisptr *Args) Run(instance *manager.Instance) {
	go instance.Run()
}

// Args ...
type Args struct {
	FileName string `xml:"FileName"`
}

func init() {
	manager.ManagerRegister["BasicControl"] = func(config []byte) func(instance *manager.Instance) {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func(instance *manager.Instance) {
				return
			}
		}
		return args.Run
	}
}
