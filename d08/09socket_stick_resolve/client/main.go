package main

import (
	"fmt"
	proto "gostudy/day08/09socket_stick_resolve/protocol"
	"net"
)

// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		b, _ := proto.Encode(msg)
		//_, _ = conn.Write([]byte(msg))
		_, _ = conn.Write(b)
	}
}
