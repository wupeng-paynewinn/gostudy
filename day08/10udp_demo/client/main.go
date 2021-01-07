package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		msg, _ := reader.ReadString('\n')
		_, err = socket.Write([]byte(msg)) // 发送数据
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}
		reply := make([]byte, 4096)
		n, remoteAddr, err := socket.ReadFromUDP(reply) // 接收数据
		if err != nil {
			fmt.Println("接收数据失败，err:", err)
			return
		}
		fmt.Printf("收到回复:%v addr:%v count:%v\n", string(reply[:n]), remoteAddr, n)
	}

}
