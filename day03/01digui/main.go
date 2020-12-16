package main

import "fmt"

//递归：函数自己调用自己
//递归要有一个明确的退出条件


//计算n的阶乘
func f(n  uint64) uint64 {
	if n<= 1{
		return 1
	}
	return n * f(n-1)
}

//n个台阶，一次可以走一步，也可以走两步，有多少种走法
func t(n uint64)  uint64{
	if n == 1{
		return 1
	}
	if n ==2 {
		return 2
	}
	return t(n-1) + t(n-2)
}


func main()  {
	fmt.Println(f(7))
	fmt.Println(t(3))
}
