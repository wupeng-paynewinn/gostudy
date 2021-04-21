package main

import (
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

func add() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
}
