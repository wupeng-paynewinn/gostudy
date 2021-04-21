package main

import (
	"fmt"
	"os"
)

// 1、文件对象的类型
// 2、获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./file.log.err")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	// 1、文件对象的类型
	fmt.Printf("%T\n", fileObj)
	// 2、获取文件对象的详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("get file failed,err:", err)
		return
	}
	fmt.Printf("File size:%dB\n", fileInfo.Size())
	fmt.Printf("File name:%s\n", fileInfo.Name())
}
