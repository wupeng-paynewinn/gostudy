package main

import "fmt"

//给自定义类型添加方法
type myInt int

//func (i int) hello() {   //不能给内置类型或者别的包类型定义方法。只能给自己的包里的自定义类型的定义方法
func (m myInt) hello() {
	fmt.Println("it's int")
}

func main() {
	//
	m := myInt(100)   //int8(100)
	m.hello()
}
