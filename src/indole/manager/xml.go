package manager

import (
	"encoding/xml"
	"io"
)

// Manager ...
type Manager struct {
	Plugin     []*Plugin     `xml:"Plugin"`
	Connection []*Connection `xml:"Connection"`
	Control    *Plugin       `xml:"Control"`
}

// Plugin ...
type Plugin struct {
	Name string `xml:"Name,attr"`
	Args string `xml:",innerxml"`
}

// Connection ...
type Connection struct {
	X    int `xml:"x,attr"`
	Y    int `xml:"y,attr"`
	Size int `xml:"size,attr"`
}

// Run ...
func (thisptr *Manager) Run() {
	f := make([]func() io.ReadWriteCloser, len(thisptr.Plugin))
	for i, v := range thisptr.Plugin {
		if gen, ok := PluginRegister[v.Name]; ok {
			if config, err := xml.Marshal(v); err == nil {
				f[i] = gen(config)
			}
		}
	}
	if ctrl, ok := ManagerRegister[thisptr.Control.Name]; ok {
		if config, err := xml.Marshal(thisptr.Plugin); err == nil {
			ctrl(config)(&Instance{
				F: f,
				E: thisptr.Connection,
			})
		}
	}
}

// PluginRegister ...
var PluginRegister = map[string]func(config []byte) func() io.ReadWriteCloser{}

// ManagerRegister ...
var ManagerRegister = map[string]func(config []byte) func(instance *Instance){}
