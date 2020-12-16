package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func log1() {
	fileObj, err := os.OpenFile("./test.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)
	}
}

/*
日志库：
1、支持往不同地方输出日志
2、日志分级别
	1、debug
	2、Trace
	3、Warning
	4、Info
	5、Error
	6、Fatal
3、日志要支持开关控制
4、完整的记录日志要包含有时间、行号、文件名、日志级别、日志信息
5、日志文件要切割
*/
func main() {
	log1()
}
