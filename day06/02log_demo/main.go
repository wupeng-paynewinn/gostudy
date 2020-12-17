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

func main() {
	log1()
}
