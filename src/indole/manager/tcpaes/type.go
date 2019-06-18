package tcpaes

import (
	"indole/core"
	"indole/plugin/aesdec"
	"indole/plugin/aesenc"
	"indole/plugin/tcp"
	"log"
	"net"
)

// TCPAES ...
type TCPAES struct {
	listener net.Listener
	server   bool
	bufsize  int
	AESENC   *aesenc.Args
	AESDEC   *aesdec.Args
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
			e := aesenc.New(thisptr.AESENC)
			defer e.Close()
			d := aesdec.New(thisptr.AESDEC)
			defer d.Close()
			c := make(chan struct{}, 4)

			if thisptr.server {
				e, d = d, e
			}

			go core.Core(x, e, thisptr.bufsize, c)
			go core.Core(e, y, thisptr.bufsize, c)
			go core.Core(y, d, thisptr.bufsize, c)
			go core.Core(d, x, thisptr.bufsize, c)

			select {
			case <-c:
			}
			log.Println("manager", "toy", "Run", "DONE")
		}()
	}
}
