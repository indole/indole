package tcpaes

import (
	"indole/manager"
	"indole/plugin/aesdec"
	"indole/plugin/aesenc"
	"indole/plugin/plain"
	"indole/plugin/tcp"
	"io"
	"log"
	"net"
)

// Build ...
func Build(args *Args) manager.Manager {
	log.Println("[manager]", "[tcpaes]", "[New]", "args:", args)
	listener, err := net.Listen(args.Network, args.Address)
	if err != nil {
		log.Println("[manager]", "[tcpaes]", "[New]", "err:", err)
		return nil
	}
	return &TCPAES{
		listener: listener,
		bufsize:  args.BufSize,
		encode:   args.Encode.extract(),
		decode:   args.Decode.extract(),
		TCP:      args.TCP,
	}
}

// Args ...
type Args struct {
	Network string    `xml:"network,attr"`
	Address string    `xml:"address,attr"`
	BufSize int       `xml:"bufsize,attr"`
	Encode  coder     `xml:"encode"`
	Decode  coder     `xml:"decode"`
	TCP     *tcp.Args `xml:"tcp"`
}

type coder struct {
	AESENC []*aesenc.Args `xml:"aesenc"`
	AESDEC []*aesdec.Args `xml:"aesdec"`
	PLAIN  []*plain.Args  `xml:"plain"`
}

func (thisptr *coder) extract() func() (ret []io.ReadWriteCloser) {
	return func() (ret []io.ReadWriteCloser) {
		for _, v := range thisptr.AESENC {
			ret = append(ret, aesenc.Build(v))
		}
		for _, v := range thisptr.AESDEC {
			ret = append(ret, aesdec.Build(v))
		}
		for _, v := range thisptr.PLAIN {
			ret = append(ret, plain.Build(v))
		}
		return
	}
}
