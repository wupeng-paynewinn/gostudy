package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan()  {
	var s string
	fmt.Print("scan:请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("scan:你输入的内容是：%s\n",s)
}

func useBufio() {
	var s string
	//os.Stdin 终端的标准输入
	reader := bufio.NewReader(os.Stdin)  //newReader的参数就是接口类型
	fmt.Print("reader:请输入内容：")
	s,_ = reader.ReadString('\n')
	fmt.Printf("reader:你输入的内容是：%v\n",s)
}
func main() {
	useScan()
	useBufio()
}
