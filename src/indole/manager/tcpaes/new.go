package tcpaes

import (
	"indole/manager"
	"indole/plugin/aesdec"
	"indole/plugin/aesenc"
	"indole/plugin/hexdecodestream"
	"indole/plugin/hexencodestream"
	"indole/plugin/plainpacket"
	"indole/plugin/plainstream"
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
	AESENC          []*aesenc.Args          `xml:"aesenc"`
	AESDEC          []*aesdec.Args          `xml:"aesdec"`
	PlainStream     []*plainstream.Args     `xml:"PlainStream"`
	PlainPacket     []*plainpacket.Args     `xml:"PlainPacket"`
	HexEncodeStream []*hexencodestream.Args `xml:"HexEncodeStream"`
	HexDecodeStream []*hexdecodestream.Args `xml:"HexDecodeStream"`
}

func (thisptr *coder) extract() func() (ret []io.ReadWriteCloser) {
	return func() (ret []io.ReadWriteCloser) {
		for _, v := range thisptr.AESENC {
			ret = append(ret, aesenc.Build(v))
		}
		for _, v := range thisptr.AESDEC {
			ret = append(ret, aesdec.Build(v))
		}
		for _, v := range thisptr.PlainStream {
			ret = append(ret, plainstream.Build(v))
		}
		for _, v := range thisptr.PlainPacket {
			ret = append(ret, plainpacket.Build(v))
		}
		for _, v := range thisptr.HexEncodeStream {
			ret = append(ret, hexencodestream.Build(v))
		}
		for _, v := range thisptr.HexDecodeStream {
			ret = append(ret, hexdecodestream.Build(v))
		}
		return
	}
}
