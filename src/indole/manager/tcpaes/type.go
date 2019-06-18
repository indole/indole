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
	network  string
	address  string
	bufsize  int
	hexkey   string
	limit    uint64
	server   bool
}

// Run ...
func (thisptr *TCPAES) Run() {
	for {
		conn, err := thisptr.listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go func() {
			x := tcp.New(conn)
			defer x.Close()
			y := tcp.NewByDial(thisptr.network, thisptr.address)
			defer y.Close()
			e := aesenc.NewBySizeHexKey(thisptr.bufsize, thisptr.hexkey)
			defer e.Close()
			d := aesdec.NewBySizeHexKeyLimit(int(thisptr.limit), thisptr.hexkey)
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
