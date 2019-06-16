package core

import (
	"io"
	"log"
)

// Core ...
func Core(x, y io.ReadWriteCloser, bufsize int, done chan struct{}) {
	defer func() {
		recover()
		done <- struct{}{}
	}()
	buf := make([]byte, bufsize)
	for {
		n, err := x.Read(buf)
		if err != nil {
			log.Println("core", "Core", "x.Read(buf)", err)
			return
		}
		_, err = y.Write(buf[:n])
		if err != nil {
			log.Println("core", "Core", "y.Write(buf[:n])", err)
			return
		}
	}
}
