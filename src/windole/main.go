package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"

	"github.com/lxn/walk"
)

const (
	iconid  = 3
	tooltip = "Windole"
)

func main() {
	mw, err := walk.NewMainWindow()
	if err != nil {
		log.Fatal(err)
	}

	icon, err := walk.NewIconFromResourceId(iconid)
	if err != nil {
		log.Fatal(err)
	}

	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()

	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}

	if err := ni.SetToolTip(tooltip); err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	cmd := fmt.Sprintf("start http://localhost:%v/", listener.Addr().(*net.TCPAddr).Port)

	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		exec.Command("cmd", "/c", cmd).Run()
	})

	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	if err := ni.ShowInfo("Windole", "windole is running ..."); err != nil {
		log.Fatal(err)
	}

	go http.Serve(listener, nil)

	mw.Run()
}
