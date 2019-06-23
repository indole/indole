package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"

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
		}),
		register("Close System Proxy", func() {
			exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings", "/v", "ProxyEnable", "/t", "REG_DWORD", "/d", "0", "/f").Run()
		}),
		register("Windole Setting", func() {

		}),
		register("Exit Windole", func() {
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

func launch() {
	defer func() {
		recover()
	}()

	if cmd != nil {
		clean()
	}

	bs, err := ioutil.ReadFile("config.xml")
	if err != nil {
		panic(err)
	}
	cmd = exec.Command("indole.exe")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	cmd.Start()
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	stdin.Write(bs)
	stdin.Close()
}

func clean() {
	defer func() {
		recover()
	}()
	cmd.Process.Signal(os.Interrupt)
	cmd.Process.Wait()
	cmd = nil
}

var cmd *exec.Cmd
