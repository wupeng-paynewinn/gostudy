package main

import "fmt"

// select
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次:", i)
		select {
		case x := <-ch:
			fmt.Printf("取值:%d\n", x)
		case ch <- i:
			fmt.Printf("存值:%d\n", i)
		}
	}
}
