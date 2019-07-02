package main

import (
	"encoding/xml"
	"indole/manager"
	_ "indole/manager/basiccontrol"
	_ "indole/manager/tcpcontrol"
	_ "indole/plugin/aesdecodepacket"
	_ "indole/plugin/aesencodepacket"
	_ "indole/plugin/createfileinterface"
	_ "indole/plugin/dynamicudpinterface"
	_ "indole/plugin/externalprocessinterface"
	_ "indole/plugin/fileloginterface"
	_ "indole/plugin/hexdecodestream"
	_ "indole/plugin/hexencodestream"
	_ "indole/plugin/openfileinterface"
	_ "indole/plugin/packettostream"
	_ "indole/plugin/packettostreamwithaes"
	_ "indole/plugin/plainpacket"
	_ "indole/plugin/plainstream"
	_ "indole/plugin/streamtopacket"
	_ "indole/plugin/streamtopacketwithaes"
	_ "indole/plugin/tcpinterface"
	_ "indole/plugin/tuninterface"
	_ "indole/plugin/udpinterface"
	"log"
	"os"
	"os/signal"
)

func main() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	for _, v := range config.Manager {
		v.Run()
	}
	select {
	case <-channel:
	}
}

var config = &struct {
	Manager []*manager.Manager `xml:"Manager"`
}{}

func init() {
	decoder := xml.NewDecoder(os.Stdin)
	err := decoder.Decode(config)
	if err != nil {
		log.Fatalln("[ERRO]", "[main]", "[init]", "decoder.Decode(config)", err)
	}
}
