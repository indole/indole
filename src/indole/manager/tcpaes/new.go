package tcpaes

import (
	"indole/manager"
	"indole/plugin/aesdec"
	"indole/plugin/aesenc"
	"indole/plugin/tcp"
	"log"
	"net"
)

// NewByArgs ...
func NewByArgs(args *Args) manager.Manager {
	log.Println("manager", "tcpaes", "NewByArgs", args)
	listener, err := net.Listen(args.Network, args.Address)
	if err != nil {
		log.Fatalln("manager", "tcpaes", "NewByArgs", "net.Listen(args.SrcNetwork, args.SrcAddress)", err)
	}
	return &TCPAES{
		listener: listener,
		server:   args.Server,
		bufsize:  args.BufSize,
		AESENC:   args.AESENC,
		AESDEC:   args.AESDEC,
		TCP:      args.TCP,
	}
}

// Args ...
type Args struct {
	Network string       `xml:"network,attr"`
	Address string       `xml:"address,attr"`
	Server  bool         `xml:"server,attr"`
	BufSize int          `xml:"bufsize,attr"`
	AESENC  *aesenc.Args `xml:"aesenc"`
	AESDEC  *aesdec.Args `xml:"aesdec"`
	TCP     *tcp.Args    `xml:"tcp"`
}
