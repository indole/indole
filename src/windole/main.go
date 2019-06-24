package main

import (
	"flag"
	"log"
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

	launch()

	actions := []*walk.Action{
		register("Open System Proxy", func() {
			exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "1", "/f").Run()
			exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyServer", "/t", "REG_SZ", "/d", proxy, "/f").Run()
		}),
		register("Close System Proxy", func() {
			exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "0", "/f").Run()
		}),
		register("Windole Setting", func() {
			ni.ShowInfo("Windole", "TODO")
		}),
		register("Exit Windole", func() {
			exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "0", "/f").Run()
			clean()
			walk.App().Exit(0)
		}),
	}
	for _, v := range actions {
		if err := ni.ContextMenu().Actions().Add(v); err != nil {
			log.Fatal(err)
		}
	}

	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	if err := ni.ShowInfo("Windole", "windole is running ..."); err != nil {
		log.Fatal(err)
	}

	mw.Run()
}

func register(name string, f func()) *walk.Action {
	action := walk.NewAction()
	if err := action.SetText(name); err != nil {
		log.Fatal(err)
	}
	action.Triggered().Attach(f)
	return action
}

var (
	config string
	indole string
	proxy  string
)

func init() {
	flag.StringVar(&config, "config", "config.xml", "config path")
	flag.StringVar(&indole, "indole", "indole.exe", "indole path")
	flag.StringVar(&proxy, "proxy", "127.0.0.1:3023", "proxy")
	flag.Parse()
}
