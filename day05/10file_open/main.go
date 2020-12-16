package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFormFile1()  {
	//循环读取文件，有点low
	fileObj,err := os.Open("./main.go")
	if err != nil{
		fmt.Println("error:",err)
		return
	}

	//关闭文件
	defer fileObj.Close()

	//var tmp = make([]byte,128)
	var tmp = [128]byte{}   //数组
	for{
		n,err := fileObj.Read(tmp[:])
		if err != nil{
			fmt.Printf("err:%v\n",err)
			return
		}
		if err == io.EOF{
			fmt.Println("读完")
			return
		}
		fmt.Printf("读了%v个字节\n",n)
		fmt.Println(string(tmp[:n]))
		if n < 128{
			return
		}
	}
}

func readFormFilebyBufio()  {
	fileObj,err := os.Open("./main.go")
	if err != nil{
		fmt.Println("打开文件error:",err)
		return
	}

	//关闭文件
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for{
		line,err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("读文件err:",err)
			return
		}
		if err == io.EOF{
			fmt.Println("读完了")
			return
		}
		fmt.Print(line)
	}

}
func readFromFileByIoutil()  {
	fileObj,err := ioutil.ReadFile("./main.go")
	if err != nil{
		fmt.Println("打开文件error:",err)
		return
	}
	fmt.Println(string(fileObj))
}
func main()  {
	//readFormFile1()
	//readFormFilebyBufio()
	readFromFileByIoutil()
}
