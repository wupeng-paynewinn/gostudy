package main

import "fmt"

func main(){
	//数组的第二种声明方法
	a1 	:=	[...]int{1,2,3,6,7,8}
	//切片
	s1 	:= 	a1 	[:]					//切片s1
	s2	:= 	a1	[2:]
	s3	:=	a1	[:4]
	//切片之后再切片
	s4 	:= 	s3	[1:]
	s5	:=	a1	[2:3]

	fmt.Println("s1",s1)
	fmt.Println("s2",s2)
	fmt.Println("s3",s3)
	fmt.Println("s4",s4)
	fmt.Println("s5",s5)

	var s6 []int    //没有分配内存 == nil
	fmt.Println("s6",s6)
	s6 = []int{1,2,3}
	fmt.Println("s6",s6)


	//make初始化，分配内存
	s7 := make([]int,2,4)
	fmt.Println("s7:",s7)


//切片不存值，它就像一个框，去底层数组取值

//切片：指针、长度、容量
//切片的扩容策略
/*
   	1、长度小于1024，直接2倍
	2、大于1024，按照1.25倍扩容
	3、申请容量大于旧的容量的2倍，直接扩容至新申请的容量
	4、具体存储的值的类型不同，扩容策略也有一定的不同
 */

	//s8 := make([]int,1,20)
	var s8 []int
	s8 = append(s8,1)  //append 自动初始化切片
	fmt.Println("s8",s8)


	s9 := []int{1,2,3}    //s9:[1 2 3]
	s10 := s9   //s10:[1 2 3]
	s11 := make([]int,3,3)


	copy(s11,s9)    //s11:[1 2 3]   把s9的值拷贝给s11

	fmt.Println("s9",s9)
	fmt.Println("s10:",s10)
	s10[1] = 200
	fmt.Println("s10",s10)  //s10:[1 200 3]
	fmt.Println("s9",s9)    // s9:[1 200  3]  ?

	fmt.Println("s11",s11)    // s11:[1 2 3]
}
