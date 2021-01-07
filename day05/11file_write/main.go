package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//打开文件写内容
func writeDemo1() {
	//os.O_APPEND：追加
	fileObj, err := os.OpenFile("./xx.html", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 4556) // “|”：两个有一个为1就为1
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("寒江孤影，\n江湖故人，\n"))
	//writeString
	fileObj.WriteString("相逢何必曾相识。\n")
	fileObj.Close()
}

func writeDemo2() {
	//os.O_TRUNC：清空
	fileObj, err := os.OpenFile("./xx.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 4557) // “|”：两个有一个为1就为1
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	//write
	fileObj.Write([]byte("自从学会了\n换行，我发现\n"))
	//writeString
	fileObj.WriteString("我会写诗了\n")
	fileObj.Close()

}

func writeDemo3() {
	//os.O_TRUNC：清空
	fileObj, err := os.OpenFile("./xx.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 4557) // “|”：两个有一个为1就为1
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("Hello monkey!")
	wr.Flush() //缓存到文件中
}

//往两个字符之中插入一个字符
func writeDemo4() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	//没有办法直接再文件中间插入内容，所以要借助一个临时文件：打开或者创建临时文件
	tmpFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed,err:%v\n", err)
		return
	}
	//defer tmpFile.Close()

	//读取文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed,err:%v\n", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])

	//再写入要插入的内容
	var s []byte
	s = []byte{'c'}

	tmpFile.Write(s)

	//紧接着把源文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}

	//源文件后续的写入临时文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")

	//直接插入会覆盖当前位置的字符
	//fileObj.Seek(1,0) //光标向右移动一位
	//var s[]byte
	//s = []byte{'c'}
	//
	//fileObj.Write(s)

}
func main() {
	//writeDemo1()
	//writeDemo2()
	//writeDemo3()
	writeDemo4()
}
