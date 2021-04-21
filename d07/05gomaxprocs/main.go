package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}

}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}

}

func main() {
	runtime.GOMAXPROCS(2) //占用2核cpu   //默认为cpu的逻辑核心数，默认跑满整个cpu
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
