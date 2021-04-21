package main

import "fmt"

func main(){
	//函数内部没有办法申明带名字的函数
	f1 := func(m,n int) {
		fmt.Println( m+ n)
	}
	f1(1,6)


	//只是调用一次的函数 ，可以简写成立即执行函数
	func (x,y int){
		fmt.Println(x+y)
	}(10,20)

}
