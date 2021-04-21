package main

import (
	"fmt"
	"sync"
)

var b chan int //chan是一个引用类型
var wg sync.WaitGroup

//通道的操作
//  1、发送：		ch1 <-  1
//	2、接收: 	x:  <- ch1
//	3、关闭： 	close()
func noBufferChannel() {
	//不带缓冲区通道的初始化
	b = make(chan int) //通道必须使用make函数进行初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道中取到了:", x)
	}()

	b <- 10 //hang住了
	fmt.Println("10发送到通达b中")
	fmt.Println(b)
	wg.Wait()
}

func bufferChannel() {
	// 带缓冲区通道的初始化
	b = make(chan int, 1) //通道必须使用make函数进行初始化
	b <- 10               //hang住了
	fmt.Println("10发送到通达b中")
	//b <- 20
	//fmt.Println("20发送到通达b中")   //设定长度为1，接收2个数会发生死锁
	fmt.Println(b)
	x := <-b
	fmt.Println("后台goroutine从通道中取到了:", x)
	//关闭通道，不关也没关系，自动垃圾回收
	close(b)

}

func main() {
	//noBufferChannel()
	bufferChannel()
}
