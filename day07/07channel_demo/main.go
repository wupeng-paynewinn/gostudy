package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	//close(ch2)
	once.Do(func() {
		close(ch2)
	}) //确保某个操作只执行一次
}

//单向通道：ch3只能接收值，ch4只能发送值
func f3(ch3 chan<- int, ch4 <-chan int) {

}
func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for i := range b {
		fmt.Println(i)
	}

}
