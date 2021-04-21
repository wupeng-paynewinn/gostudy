package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 500000; i++ {
		// 对公共资源的访问，都需要加锁
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
