package main

import (
	"encoding/xml"
	"indole/manager/tcpaes"
	"indole/manager/udptun"
	"log"
	"os"
	"os/signal"
)

func main() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	for _, v := range config.TCPAES {
		manager := tcpaes.Build(v)
		go manager.Run()
	}
	for _, v := range config.UDPTUN {
		manager := udptun.Build(v)
		go manager.Run()
	}

	select {
	case <-channel:
	}
}

var config = &struct {
	TCPAES []*tcpaes.Args `xml:"tcpaes"`
	UDPTUN []*udptun.Args `xml:"udptun"`
}{}

func init() {
	decoder := xml.NewDecoder(os.Stdin)
	err := decoder.Decode(config)
	if err != nil {
		log.Fatalln("main", "init", "decoder.Decode(config)", err)
	}
}
