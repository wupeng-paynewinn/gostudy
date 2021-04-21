package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen udp err:", err)
	}
	defer conn.Close()
	var data [1024]byte
	for {
		//接收数据
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			return
		}
		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		_, _ = conn.WriteToUDP([]byte(reply), addr)
	}

}
