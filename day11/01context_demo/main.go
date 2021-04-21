package main

import (
	"fmt"
	"sync"
	"time"
)

//主goroutine通知子goroutine退出：

var wg sync.WaitGroup

//第一种 采用全局变量的方式通知
var notify bool

var exitChan = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("打印")
		time.Sleep(time.Millisecond * 500)
		//方式一
		if notify {
			break
		}
		//方式二
		select {
		case <-exitChan:
			//break挑不出for循环，所以加入tag，指定跳出
			break LOOP
		default:
		}

	}
}
func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 如何优雅的通知goroutine退出
	//方式一
	notify = true
	//方式二
	exitChan <- true
	wg.Wait()

}
