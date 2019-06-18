package toy

import (
	"indole/core"
	"indole/plugin/plain"
	"indole/plugin/tcp"
	"log"
	"net"
)

// Toy ...
type Toy struct {
	listener net.Listener
	network  string
	address  string
	bufsize  int
}

// Run ...
func (thisptr *Toy) Run() {
	for {
		conn, err := thisptr.listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go func() {
			x := tcp.NewByConn(conn)
			defer x.Close()
			y := tcp.New(&tcp.Args{
				Network: thisptr.network,
				Address: thisptr.address,
			})
			defer y.Close()
			e := plain.NewBySize(thisptr.bufsize)
			defer e.Close()
			d := plain.NewBySize(thisptr.bufsize)
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
