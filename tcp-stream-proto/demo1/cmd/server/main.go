package main

import (
	"fmt"
	"net"

	"github.com/bigwhite/tcp-stream-proto/demo1/pkg/frame"
	"github.com/bigwhite/tcp-stream-proto/demo1/pkg/packet"
)

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()

	for {
		// read from the connection
		framePayLoad, err := frameCodec.Decode(c)
		if err != nil {
			panic(err)
		}

		p, err := packet.Decode(framePayLoad)
		var ackFramePayload []byte

		// do something with p
		switch p.(type) {
		case *packet.Submit:
			submit := p.(*packet.Submit)
			submitAck := &packet.SubmitAck{
				ID:     submit.ID,
				Result: 0,
			}
			ackFramePayload, err = packet.Encode(submitAck)
			if err != nil {
				panic(err)
			}
		default:
			//...
		}
		// write to the connection
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			panic(err)
		}
	}

}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
