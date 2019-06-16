package tcpaess

import (
	"indole/manager"
	"log"
	"net"
)

// New ...
func New(listener net.Listener, network, address string, bufsize int, hexkey string, limit uint64) manager.Manager {
	return &TCPAESS{
		listener: listener,
		network:  network,
		address:  address,
		bufsize:  bufsize,
		hexkey:   hexkey,
		limit:    limit,
	}
}

// NewByArgs ...
func NewByArgs(args *Args) manager.Manager {
	log.Println("manager", "toy", "NewByArgs", args)
	listener, err := net.Listen(args.SrcNetwork, args.SrcAddress)
	if err != nil {
		log.Fatalln("manager", "toy", "NewByArgs", "net.Listen(args.SrcNetwork, args.SrcAddress)", err)
	}
	return New(listener, args.DstNetwork, args.DstAddress, args.BufSize, args.HexKey, args.Limit)
}

// Args ...
type Args struct {
	SrcNetwork string `xml:"src_network,attr"`
	SrcAddress string `xml:"src_address,attr"`
	DstNetwork string `xml:"dst_network,attr"`
	DstAddress string `xml:"dst_address,attr"`
	BufSize    int    `xml:"buf_size,attr"`
	HexKey     string `xml:"hex_key,attr"`
	Limit      uint64 `xml:"limit,attr"`
}
