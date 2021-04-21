package main

import "fmt"

//接口是一种类型，一种抽象的类型
//引入接口的实例

type cat struct {}
type dog struct {}
type person struct {}

type speaker interface {
	//只要实现了speak方法的变量，都能当成speak类型
	speak()
}

func (c cat)speak()  {
	fmt.Println("喵喵喵~")
}
func (d dog)speak()  {
	fmt.Println("汪汪汪~")
}

func (p person) speak()  {
	fmt.Println("啊啊啊~")
}

func da(x speaker)  {
	//接收一个参数，传进来什么，我就打什么
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker
	ss = c1
	fmt.Println(ss)

}
