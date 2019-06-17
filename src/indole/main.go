package main

import (
	"encoding/xml"
	"indole/manager/tcpaes"
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
	for _, v := range config.TCPAES {
		manager := tcpaes.NewByArgs(v)
		go manager.Run()
	}

	select {
	case <-channel:
	}
}

var config = &struct {
	Toy    []*toy.Args    `xml:"toy"`
	TCPAES []*tcpaes.Args `xml:"tcpaes"`
}{}

func init() {
	decoder := xml.NewDecoder(os.Stdin)
	err := decoder.Decode(config)
	if err != nil {
		log.Fatalln("main", "init", "decoder.Decode(config)", err)
	}
}
