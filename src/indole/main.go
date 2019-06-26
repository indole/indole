package main

import (
	"encoding/xml"
	"indole/manager"
	_ "indole/manager/basiccontrol"
	_ "indole/manager/tcpcontrol"
	_ "indole/plugin/createfileinterface"
	_ "indole/plugin/openfileinterface"
	_ "indole/plugin/tcpinterface"
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
		log.Fatalln("main", "init", "decoder.Decode(config)", err)
	}
}
