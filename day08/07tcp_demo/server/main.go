package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// tcp/server/main.go

// TCP server端
var wg sync.WaitGroup
var goroutineCount = 100

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		//接收数据
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		// 发送数据
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, _ = conn.Write([]byte(inputInfo)) // 发送数据
	}
	wg.Done()
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		wg.Add(goroutineCount)
		for i := 0; i < goroutineCount; i++ {
			go process(conn) // 启动一个goroutine处理连接
		}
		wg.Wait()
	}
}
