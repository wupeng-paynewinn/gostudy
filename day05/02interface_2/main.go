package main

import "fmt"

/*
接口的定义格式：
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
 */

type animal interface {
	move()
	eat(s string)
}

type cat struct {
	name string
	feet int8
}

func (c cat)move()  {
	fmt.Println("猫动")
}
func (c  cat)eat(food string)  {
	fmt.Printf("吃%s!\n",food)
}


type chicken struct {
	feet int8
}

func (c chicken)move()  {
	fmt.Println("鸡动")
}
func (c  chicken)eat(food string)  {
	fmt.Printf("吃%s！\n",food)
}


func main() {
	var a1 animal //定义一个接口类型的变量
	fmt.Printf("%T\n",a1)  //类型：nil （动态类型）

	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)


	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("a1 type:%T",a1)  //a1 type:main.chicken
}
