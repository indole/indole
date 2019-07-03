package manager

import (
	"indole/utils"
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
				defer utils.Recover("[WARN]", "[manager]", "[Instance]", "[Run]", "[Close]")
				v.Close()
			}(v)
		}
	}()
	c := make(chan struct{}, len(thisptr.E))
	for _, v := range thisptr.E {
		go func(x, y io.ReadWriteCloser, size int) {
			defer func() {
				c <- struct{}{}
			}()
			defer utils.Recover("[WARN]", "[manager]", "[Instance]", "[Run]")
			buf := make([]byte, size)
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
		}(vs[v.X], vs[v.Y], v.Size)
	}
	select {
	case <-c:
	}
}
