package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	//当并发多了之后（大于20）执行就会报fatal error: concurrent map writes错误
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			// Go语言中内置的map不是并发安全的
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
