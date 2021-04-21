package main

import "fmt"

//函数
//函数的定义

//函数是一段代码的封装

func sum(x int,y int)(ret int){
	return x+y
}


//没有返回值
func f1 (x int ,y  int)(){
	fmt.Println(x+y)
}

//没有参数没有返回值
func f2(){
	fmt.Println("f2")
}

//没有参数但有返回值
func f3 ()int{
	fmt.Println()
	return 3
}

//返回值可以命名也可以不命名
func f4(x int ,y int)(ret int){
	return x+y
	
}
//多个返回值
func f5()(int,string)  {
	return 1,"开始"
}

// 参数的类型简写
func f6(x,y,z int,m,n string,i,j bool) (sum int,name string ,isOk bool) {

	return x+y+z, m+n,i&&j
}


//可变长参数
//可变长参数必须放在函数参数的最后
func f7(x string,y ...int)  {
	fmt.Println(x)
	fmt.Println(y) // y的类型是切片 []int
}


//go语言没有默认参数这个概念

func main(){
	r := sum(1,2)
	fmt.Println(r)
}
