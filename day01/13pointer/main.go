package main

import "fmt"

func main(){
	// Potential nil pointer dereference
	var a1  *int    //空指针<nil>
	fmt.Println(a1)
	//*a = 100
	//fmt.Println(*a)

	//&：取地址；   *：根据地址取值
	//go里面的指针只能读，不能修改

	/*
	make和new的区别:
	1.make和new都是用来申请内存的
	2.new很少用，一般用来给基本数据类型申请内存，string、 int,返回的是对应类型的指针(*string、*int)。
	3.make是用来给slice、map、 chan申请内存的，make函数返回的的是对应的这三个类型本身
	*/

	//使用new()创建一个指针
	var a  =  new(int)
	fmt.Println(a)

	fmt.Println(*a)

	*a =100
	fmt.Println(*a)

	//案例：
	addr := "ths"
	addrP := &addr      //内存地址
	fmt.Println("addrP",addrP)
	fmt.Printf("%T\n",addrP)

	addrV := *addrP     	//根据内存地址找值
	fmt.Println("addrV:",addrV)


}
