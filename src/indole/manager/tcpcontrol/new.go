package tcpcontrol

import (
	"encoding/xml"
	"indole/manager"
	"indole/plugin/tcpinterfacebyconn"
	"io"
	"log"
	"net"
)

// Run ...
func (thisptr *Args) Run(instance *manager.Instance) {
	listener, err := net.Listen(thisptr.Network, thisptr.Address)
	if err != nil {
		log.Println("[manager]", "[tcpcontrol]", "[Run]", "err:", err)
	}
	conns := make(chan net.Conn, 0)
	instance.E = append(instance.E, &manager.Connection{
		X:    len(instance.F),
		Y:    thisptr.In,
		Size: thisptr.Size,
	})
	instance.E = append(instance.E, &manager.Connection{
		X:    thisptr.Out,
		Y:    len(instance.F),
		Size: thisptr.Size,
	})
	instance.F = append(instance.F, func() io.ReadWriteCloser {
		return (&tcpinterfacebyconn.Args{
			Conn: <-conns,
		}).Build()
	})
	go func() {
		for {
			if conn, err := listener.Accept(); err == nil {
				go func() {
					conns <- conn
				}()
				go instance.Run()
			}
		}
	}()
}

// Args ...
type Args struct {
	Network string `xml:"Network"`
	Address string `xml:"Address"`
	In      int    `xml:"In"`
	Out     int    `xml:"Out"`
	Size    int    `xml:"Size"`
}

func init() {
	manager.ManagerRegister["TCPControl"] = func(config []byte) func(instance *manager.Instance) {
		args := &Args{}
		if err := xml.Unmarshal(config, args); err != nil {
			return func(instance *manager.Instance) {
				return
			}
		}
		return args.Run
	}
}
