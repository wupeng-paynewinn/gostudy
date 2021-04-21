package main

import (
	"fmt"
)

//闭包：应用场景
//底层原理：
//1.函数可以作为返回值
//2、函数内部查找变量的顺序，现在自己的内部找，找不到往外层找
//闭包 = 函数 +外部变量的引用

func f1 (f func()){
	fmt.Println("is f1")
	f()
}

func f2(x,y int){
	fmt.Println("is f2")
	fmt.Println( x+y)
}


//要求：实现 f1(f2)

//定义一个函数对f2进行包装
func f3(f func(int,int),x,y int) func(){
	return func() {
		fmt.Println(x+y)
		//调用一下f2
		f(x,y)
	}
}

func main(){
	f1(f3(f2,100,200))

}
