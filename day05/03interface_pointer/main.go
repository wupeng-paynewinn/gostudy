package main

import "fmt"

/*
使用值接收者实现接口与使用指针接收者实现接口的区别？
1、使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
2、使用指针接收者实现接口只能存结构体指针类型
 */

type animal interface {
	move()
	eat(s string)
}

type cat struct {
	name string
	feet int8
}


//使用指针接收者实现接口的所有方法
func (c *cat)move()  {
	fmt.Println("猫动")
}
func (c  *cat)eat(food string)  {
	fmt.Printf("吃%s!\n",food)
}



func main() {
	var a1 animal //定义一个接口类型的变量


	c1 := cat{"淘气", 4}  //cat
	c2 := &cat{"骄傲脸",4} //*cat
	a1 = &c1     //实现animal这个接口的是cat的指针类型
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)


}
