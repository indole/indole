package manager

import (
	"io"
)

// Instance ...
type Instance struct {
	F []func() io.ReadWriteCloser
	E []*Connection
}

// Run ...
func (thisptr *Instance) Run() {
	vs := make([]io.ReadWriteCloser, len(thisptr.F))
	for i, v := range thisptr.F {
		vs[i] = v()
	}
	defer func() {
		for _, v := range vs {
			func(v io.ReadWriteCloser) {
				defer func() {
					recover()
				}()
				v.Close()
			}(v)
		}
	}()
	c := make(chan struct{}, len(thisptr.E))
	for _, v := range thisptr.E {
		go Core(vs[v.X], vs[v.Y], v.Size, c)
	}
	select {
	case <-c:
	}
}

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
