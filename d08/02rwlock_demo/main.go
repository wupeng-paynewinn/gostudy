package main

// 读写互斥锁

import (
	"fmt"
	"sync"
	"time"
)

var (
	x  = 0
	wg sync.WaitGroup
	//lock sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	defer wg.Done()
	rwlock.Lock()
	defer rwlock.Unlock()
	x = x + 1
	fmt.Println("写入", x)
	time.Sleep(time.Millisecond * 10)
}

func read() {
	defer wg.Done()
	rwlock.RLock()
	defer rwlock.RUnlock()
	fmt.Println("读取", x)
	time.Sleep(time.Millisecond * 1)
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 10000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
