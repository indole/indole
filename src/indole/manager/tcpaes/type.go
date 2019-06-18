package tcpaes

import (
	"indole/core"
	"indole/plugin/tcp"
	"io"
	"log"
	"net"
)

// TCPAES ...
type TCPAES struct {
	listener net.Listener
	bufsize  int
	encode   func() []io.ReadWriteCloser
	decode   func() []io.ReadWriteCloser
	core     func(x io.ReadWriteCloser)
	TCP      *tcp.Args
}

// Run ...
func (thisptr *TCPAES) Run() {
	for {
		conn, err := thisptr.listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go func() {
			x := tcp.NewByConn(conn)
			defer x.Close()
			y := tcp.New(thisptr.TCP)
			defer y.Close()
			e := thisptr.encode()
			defer func() {
				for _, v := range e {
					v.Close()
				}
			}()
			d := thisptr.decode()
			defer func() {
				for _, v := range d {
					v.Close()
				}
			}()
			c := make(chan struct{}, len(e)+len(d)+2)

			{
				w := x
				for _, v := range e {
					go core.Core(w, v, thisptr.bufsize, c)
					w = v
				}
				go core.Core(w, y, thisptr.bufsize, c)
			}
			{
				w := x
				for _, v := range d {
					go core.Core(v, w, thisptr.bufsize, c)
					w = v
				}
				go core.Core(y, w, thisptr.bufsize, c)
			}

			select {
			case <-c:
			}
			log.Println("manager", "toy", "Run", "DONE")
		}()
	}
}
