package tcpaesc

import (
	"indole/core"
	"indole/plugin/aesdec"
	"indole/plugin/aesenc"
	"indole/plugin/tcp"
	"log"
	"net"
)

// TCPAESC ...
type TCPAESC struct {
	listener net.Listener
	network  string
	address  string
	bufsize  int
	hexkey   string
	limit    uint64
}

// Run ...
func (thisptr *TCPAESC) Run() {
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
			d := aesdec.NewBySizeHexKeyLimit(thisptr.bufsize, thisptr.hexkey, thisptr.limit)
			defer d.Close()
			c := make(chan struct{}, 4)

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
