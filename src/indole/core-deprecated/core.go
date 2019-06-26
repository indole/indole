package core

import (
	"io"
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
			return
		}
		_, err = y.Write(buf[:n])
		if err != nil {
			return
		}
	}
}
