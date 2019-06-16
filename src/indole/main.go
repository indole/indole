package main

import (
	"encoding/xml"
	"indole/manager/tcpaesc"
	"indole/manager/tcpaess"
	"indole/manager/toy"
	"log"
	"os"
	"os/signal"
)

func main() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)

	for _, v := range config.Toy {
		manager := toy.NewByArgs(v)
		go manager.Run()
	}
	for _, v := range config.TCPAESC {
		manager := tcpaesc.NewByArgs(v)
		go manager.Run()
	}
	for _, v := range config.TCPAESS {
		manager := tcpaess.NewByArgs(v)
		go manager.Run()
	}

	select {
	case <-channel:
	}
}

var config = &struct {
	Toy     []*toy.Args     `xml:"toy"`
	TCPAESC []*tcpaesc.Args `xml:"tcpaesc"`
	TCPAESS []*tcpaess.Args `xml:"tcpaess"`
}{}

func init() {
	decoder := xml.NewDecoder(os.Stdin)
	err := decoder.Decode(config)
	if err != nil {
		log.Fatalln("main", "init", "decoder.Decode(config)", err)
	}
}
