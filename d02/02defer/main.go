package main

import (
	"fmt"
)

func deferDemo(){
	fmt.Println("Start")
	defer fmt.Println("嘿嘿嘿")   //defer 把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("哈哈哈") 		//多个defer语言按照先进后出的顺序延迟执行（先写的后执行）
	fmt.Println("End")
}


func f1()int{
	x :=5
	defer func(){
		x++ 		//修改的是x，不是返回值
	}()
	return x
}

func f2()(x int){
	defer func() {
		x++
	}()
	return 5
}

func f3()(y int){
	x :=5
	defer func() {
		x++
	}()
	return x
}

func f4()(x int){
	defer func(x int) {
		x++				//改变的是函数中x的副本
	}(x)
	return 5
}

func f5()(x int){
	defer func(x *int) {
		(*x)++				//改变的是函数中x的副本
	}(&x)
	return 5
}
func main(){
	deferDemo()
	fmt.Println(f1()) 		//5
	fmt.Println(f2())		//6
	fmt.Println(f3())		//5
	fmt.Println(f4())		//5
	fmt.Println(f5())		//6
}
