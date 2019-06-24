package main

import (
	"io/ioutil"
	"os/exec"
	"syscall"
)

func launch() {
	defer func() {
		recover()
	}()

	if cmd != nil {
		clean()
	}

	bs, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}
	cmd = exec.Command(indole)
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
	cmd.Process.Kill()
	cmd.Process.Wait()
	cmd = nil
}

var cmd *exec.Cmd
