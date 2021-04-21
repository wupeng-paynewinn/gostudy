package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
//主goroutine通知子goroutine退出：

var wg sync.WaitGroup

func f(ctx context.Context){
	defer wg.Done()
	go f2(ctx)
LOOP:
	for{
		fmt.Println("111")
		time.Sleep(time.Millisecond*500)
		//方式三 官方版本context
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}

	}
}
func f2(ctx context.Context){
	defer wg.Done()
LOOP:
	for{
		fmt.Println("2222222")
		time.Sleep(time.Millisecond*500)
		//方式三 官方版本context
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second*5)
	// 如何优雅的通知goroutine退出
	//方式三 退出全部goroutine
	cancel()
	wg.Wait()


}
