package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    //	int64
		r2 := rand.Intn(10) // 	0 <= n < 10
		fmt.Println(r1, r2)
	}

}

func f2(i int) {
	defer wg.Done() //每个被等待的线程在结束时应调用Done方法
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

func main() {
	//f1()
	for i := 0; i < 10; i++ {
		wg.Add(1) //父线程调用Add方法来设定应等待的线程的数量
		go f2(i)
	}
	wg.Wait() // 主线程里可以调用Wait方法阻塞至所有线程结束
}
