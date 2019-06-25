package main

import (
	"io/ioutil"
	"os/exec"
	"sync"
)

var cmd *exec.Cmd
var lock sync.Mutex

func restart(indole, config string) error {
	lock.Lock()
	defer lock.Unlock()
	if cmd != nil {
		cmd.Process.Kill()
	}
	bs, err := ioutil.ReadFile(config)
	if err != nil {
		return err
	}
	cmd = exec.Command(indole)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	cmd.Start()
	stdin.Write(bs)
	stdin.Close()
	return nil
}
